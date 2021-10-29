/*
 * Copyright (c) 2019 Miguel Ángel Ortuño.
 * See the LICENSE file for more information.
 */

package pgsql

import (
	"context"
	"database/sql"

	sq "github.com/Masterminds/squirrel"
	mucmodel "github.com/ortuman/jackal/model/muc"
	"github.com/ortuman/jackal/xmpp/jid"
)

type pgSQLOccupant struct {
	*pgSQLStorage
}

func newOccupant(db *sql.DB) *pgSQLOccupant {
	return &pgSQLOccupant{
		pgSQLStorage: newStorage(db),
	}
}

func (o *pgSQLOccupant) UpsertOccupant(ctx context.Context, occ *mucmodel.Occupant) error {
	return o.inTransaction(ctx, func(tx *sql.Tx) error {
		// store occupants data (except for resources)
		columns := []string{"occupant_jid", "bare_jid", "affiliation", "role"}
		values := []interface{}{occ.OccupantJID.String(), occ.BareJID.String(),
			occ.GetAffiliation(), occ.GetRole()}
		q := sq.Insert("occupants").
			Columns(columns...).
			Values(values...).
			Suffix("ON CONFLICT (occupant_jid) DO UPDATE SET affiliation = $3, role = $4")

		_, err := q.RunWith(tx).ExecContext(ctx)
		if err != nil {
			return err
		}

		//store occupants resources
		columns = []string{"occupant_jid", "resource"}
		for _, res := range occ.GetAllResources() {
			values = []interface{}{occ.OccupantJID.String(), res}
			q = sq.Insert("resources").
				Columns(columns...).
				Values(values...).
				Suffix("ON CONFLICT (occupant_jid) DO UPDATE SET resource = $2")

			_, err = q.RunWith(tx).ExecContext(ctx)
			if err != nil {
				return err
			}
		}

		return nil
	})
}

func (o *pgSQLOccupant) DeleteOccupant(ctx context.Context, occJID *jid.JID) error {
	return o.inTransaction(ctx, func(tx *sql.Tx) error {
		_, err := sq.Delete("occupants").Where(sq.Eq{"occupant_jid": occJID.String()}).
			RunWith(tx).ExecContext(ctx)
		if err != nil {
			return err
		}
		_, err = sq.Delete("resources").Where(sq.Eq{"occupant_jid": occJID.String()}).
			RunWith(tx).ExecContext(ctx)
		if err != nil {
			return err
		}
		return nil
	})
}

func (o *pgSQLOccupant) FetchOccupant(ctx context.Context, occJID *jid.JID) (*mucmodel.Occupant,
	error) {
	tx, err := o.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}

	occ, err := fetchOccupantData(ctx, tx, occJID)
	switch err {
	case nil:
	case sql.ErrNoRows:
		_ = tx.Commit()
		return nil, nil
	default:
		_ = tx.Rollback()
		return nil, err

	}

	err = fetchOccupantResources(ctx, tx, occ, occJID)
	if err != nil {
		_ = tx.Rollback()
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}
	return occ, nil
}

func fetchOccupantData(ctx context.Context, tx *sql.Tx, occJID *jid.JID) (*mucmodel.Occupant,
	error) {
	var occ *mucmodel.Occupant
	q := sq.Select("occupant_jid", "bare_jid", "affiliation", "role").
		From("occupants").
		Where(sq.Eq{"occupant_jid": occJID.String()})

	var occJIDStr, bareJIDStr, affiliation, role string
	err := q.RunWith(tx).
		QueryRowContext(ctx).
		Scan(&occJIDStr, &bareJIDStr, &affiliation, &role)
	switch err {
	case nil:
		occJIDdb, err := jid.NewWithString(occJIDStr, false)
		if err != nil {
			return nil, err
		}
		bareJID, err := jid.NewWithString(bareJIDStr, false)
		if err != nil {
			return nil, err
		}
		occ, err = mucmodel.NewOccupant(occJIDdb, bareJID)
		if err != nil {
			return nil, err
		}
		err = occ.SetAffiliation(affiliation)
		if err != nil {
			return nil, err
		}
		err = occ.SetRole(role)
		if err != nil {
			return nil, err
		}
	default:
		return nil, err
	}
	return occ, nil
}

func fetchOccupantResources(ctx context.Context, tx *sql.Tx, occ *mucmodel.Occupant,
	occJID *jid.JID) error {
	resources, err := sq.Select("occupant_jid", "resource").
		From("resources").
		Where(sq.Eq{"occupant_jid": occJID.String()}).
		RunWith(tx).QueryContext(ctx)
	if err != nil {
		return err
	}
	for resources.Next() {
		var dummy, res string
		if err = resources.Scan(&dummy, &res); err != nil {
			return err
		}
		occ.AddResource(res)
	}
	return nil
}

func (o *pgSQLOccupant) OccupantExists(ctx context.Context, occJID *jid.JID) (bool, error) {
	q := sq.Select("COUNT(*)").
		From("occupants").
		Where(sq.Eq{"occupant_jid": occJID.String()})

	var count int
	err := q.RunWith(o.db).QueryRowContext(ctx).Scan(&count)
	switch err {
	case nil:
		return count > 0, nil
	default:
		return false, err
	}
}
