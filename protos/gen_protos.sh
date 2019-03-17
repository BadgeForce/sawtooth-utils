#!/usr/bin/env bash

protoc -I=. templates/template.proto templates/payload.proto templates/transaction_receipts.proto --go_out=.
protoc -I=. credentials/credential.proto  credentials/payload.proto credentials/transaction_receipts.proto --go_out=.

mv ./github.com/BadgeForce/sawtooth-utils/protos/credentials_pb ./credentials_pb
mv ./github.com/BadgeForce/sawtooth-utils/protos/templates_pb ./templates_pb