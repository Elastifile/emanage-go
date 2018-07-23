package erun

var metabenchTemplate = `
{{with .Description}} description = {{.}}{{end}}
{{with .Scenario}} scenario = {{.}}{{end}}
{{with .QueueSizePerClient}} queue_size_per_client = {{.}}{{end}}
{{with .NrClients}} nr_clients = {{.}}{{end}}
{{with .NrDirs}} nr_dirs = {{.}}{{end}}
{{with .NrSubdirs}} nr_subdirs = {{.}}{{end}}
{{with .NrTotalFiles}} nr_total_files = {{.}}{{end}}
{{with .NrMoves}} nr_moves = {{.}}{{end}}
{{with .NrIterationsPerWorker}} nr_iterations_per_worker = {{.}}{{end}}
`
