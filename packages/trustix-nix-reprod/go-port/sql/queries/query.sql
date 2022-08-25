-- name: GetEval :one
SELECT * FROM evaluation
WHERE commit_sha = ? LIMIT 1;

-- name: CreateEval :one
INSERT INTO evaluation (commit_sha) VALUES (?) RETURNING *;

-- name: GetDerivation :one
SELECT * FROM derivation
WHERE drv = ? LIMIT 1;

-- name: CreateDerivation :one
INSERT INTO derivation (drv, system) VALUES (?, ?) RETURNING *;

-- name: CreateDerivationRefDirect :exec
INSERT INTO derivationrefdirect (drv_id, referrer_id) VALUES (?, ?);

-- name: CreateDerivationRefRecursive :exec
INSERT INTO derivationrefdirect (drv_id, referrer_id) VALUES (?, ?);
