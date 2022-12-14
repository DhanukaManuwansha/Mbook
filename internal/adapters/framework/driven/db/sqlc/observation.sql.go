// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.14.0
// source: observation.sql

package db

import (
	"context"
	"database/sql"
	"time"
)

const addObservation = `-- name: AddObservation :one
INSERT INTO "Observation" ("obs_datetime", "bp_sys",bp_dia, "pr", "rr", "temp", "notes", "patient_id", "nurse_id") 
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
RETURNING "observation_id"
`

type AddObservationParams struct {
	ObsDatetime time.Time      `json:"obs_datetime"`
	BpSys       float64        `json:"bp_sys"`
	BpDia       float64        `json:"bp_dia"`
	Pr          float64        `json:"pr"`
	Rr          float64        `json:"rr"`
	Temp        float64        `json:"temp"`
	Notes       sql.NullString `json:"notes"`
	PatientID   string         `json:"patient_id"`
	NurseID     int64          `json:"nurse_id"`
}

func (q *Queries) AddObservation(ctx context.Context, arg AddObservationParams) (int64, error) {
	row := q.db.QueryRowContext(ctx, addObservation,
		arg.ObsDatetime,
		arg.BpSys,
		arg.BpDia,
		arg.Pr,
		arg.Rr,
		arg.Temp,
		arg.Notes,
		arg.PatientID,
		arg.NurseID,
	)
	var observation_id int64
	err := row.Scan(&observation_id)
	return observation_id, err
}

const getAllBPDiaValues = `-- name: GetAllBPDiaValues :many
SELECT "bp_dia"
FROM "Observation"
WHERE "patient_id"=$1 AND "obs_datetime">= NOW() - INTERVAL '30 DAY'
`

func (q *Queries) GetAllBPDiaValues(ctx context.Context, patientID string) ([]float64, error) {
	rows, err := q.db.QueryContext(ctx, getAllBPDiaValues, patientID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []float64
	for rows.Next() {
		var bp_dia float64
		if err := rows.Scan(&bp_dia); err != nil {
			return nil, err
		}
		items = append(items, bp_dia)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getAllBPSisValues = `-- name: GetAllBPSisValues :many
SELECT "bp_sys"
FROM "Observation"
WHERE "patient_id"=$1 AND "obs_datetime">= NOW() - INTERVAL '30 DAY'
`

func (q *Queries) GetAllBPSisValues(ctx context.Context, patientID string) ([]float64, error) {
	rows, err := q.db.QueryContext(ctx, getAllBPSisValues, patientID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []float64
	for rows.Next() {
		var bp_sys float64
		if err := rows.Scan(&bp_sys); err != nil {
			return nil, err
		}
		items = append(items, bp_sys)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getAllObservations = `-- name: GetAllObservations :many
SELECT observation_id, obs_datetime, bp_sys, bp_dia, pr, rr, temp, notes, created_at, patient_id, nurse_id FROM "Observation"
WHERE "patient_id"=$1
ORDER BY "observation_id"
`

func (q *Queries) GetAllObservations(ctx context.Context, patientID string) ([]Observation, error) {
	rows, err := q.db.QueryContext(ctx, getAllObservations, patientID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Observation
	for rows.Next() {
		var i Observation
		if err := rows.Scan(
			&i.ObservationID,
			&i.ObsDatetime,
			&i.BpSys,
			&i.BpDia,
			&i.Pr,
			&i.Rr,
			&i.Temp,
			&i.Notes,
			&i.CreatedAt,
			&i.PatientID,
			&i.NurseID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getAllPRValues = `-- name: GetAllPRValues :many
SELECT "pr"
FROM "Observation"
WHERE "patient_id"=$1 AND "obs_datetime">= NOW() - INTERVAL '30 DAY'
`

func (q *Queries) GetAllPRValues(ctx context.Context, patientID string) ([]float64, error) {
	rows, err := q.db.QueryContext(ctx, getAllPRValues, patientID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []float64
	for rows.Next() {
		var pr float64
		if err := rows.Scan(&pr); err != nil {
			return nil, err
		}
		items = append(items, pr)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getAllRRValues = `-- name: GetAllRRValues :many
SELECT "rr"
FROM "Observation"
WHERE "patient_id"=$1 AND "obs_datetime">= NOW() - INTERVAL '30 DAY'
`

func (q *Queries) GetAllRRValues(ctx context.Context, patientID string) ([]float64, error) {
	rows, err := q.db.QueryContext(ctx, getAllRRValues, patientID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []float64
	for rows.Next() {
		var rr float64
		if err := rows.Scan(&rr); err != nil {
			return nil, err
		}
		items = append(items, rr)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getAllTempValues = `-- name: GetAllTempValues :many
SELECT "temp"
FROM "Observation"
WHERE "patient_id"=$1 AND "obs_datetime">= NOW() - INTERVAL '30 DAY'
`

func (q *Queries) GetAllTempValues(ctx context.Context, patientID string) ([]float64, error) {
	rows, err := q.db.QueryContext(ctx, getAllTempValues, patientID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []float64
	for rows.Next() {
		var temp float64
		if err := rows.Scan(&temp); err != nil {
			return nil, err
		}
		items = append(items, temp)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getObservationById = `-- name: GetObservationById :one
SELECT observation_id, obs_datetime, bp_sys, bp_dia, pr, rr, temp, notes, created_at, patient_id, nurse_id FROM "Observation" 
WHERE "observation_id" = $1 
LIMIT 1
`

func (q *Queries) GetObservationById(ctx context.Context, observationID int64) (Observation, error) {
	row := q.db.QueryRowContext(ctx, getObservationById, observationID)
	var i Observation
	err := row.Scan(
		&i.ObservationID,
		&i.ObsDatetime,
		&i.BpSys,
		&i.BpDia,
		&i.Pr,
		&i.Rr,
		&i.Temp,
		&i.Notes,
		&i.CreatedAt,
		&i.PatientID,
		&i.NurseID,
	)
	return i, err
}

const updateObservation = `-- name: UpdateObservation :one
UPDATE "Observation"
SET "obs_datetime" = $2, "bp_sys" = $3 ,"bp_dia" = $4, "pr" = $5 , "rr" =$6, "temp" = $7, "notes" = $8, "patient_id" = $9, "nurse_id" =$10
WHERE "observation_id" = $1
RETURNING observation_id, obs_datetime, bp_sys, bp_dia, pr, rr, temp, notes, created_at, patient_id, nurse_id
`

type UpdateObservationParams struct {
	ObservationID int64          `json:"observation_id"`
	ObsDatetime   time.Time      `json:"obs_datetime"`
	BpSys         float64        `json:"bp_sys"`
	BpDia         float64        `json:"bp_dia"`
	Pr            float64        `json:"pr"`
	Rr            float64        `json:"rr"`
	Temp          float64        `json:"temp"`
	Notes         sql.NullString `json:"notes"`
	PatientID     string         `json:"patient_id"`
	NurseID       int64          `json:"nurse_id"`
}

func (q *Queries) UpdateObservation(ctx context.Context, arg UpdateObservationParams) (Observation, error) {
	row := q.db.QueryRowContext(ctx, updateObservation,
		arg.ObservationID,
		arg.ObsDatetime,
		arg.BpSys,
		arg.BpDia,
		arg.Pr,
		arg.Rr,
		arg.Temp,
		arg.Notes,
		arg.PatientID,
		arg.NurseID,
	)
	var i Observation
	err := row.Scan(
		&i.ObservationID,
		&i.ObsDatetime,
		&i.BpSys,
		&i.BpDia,
		&i.Pr,
		&i.Rr,
		&i.Temp,
		&i.Notes,
		&i.CreatedAt,
		&i.PatientID,
		&i.NurseID,
	)
	return i, err
}
