#!/usr/bin/env bash

# assume in tendermint/tendermint

# un-comment once using for CI
# go get github.com/melekes/godoc2md

destination=$GOPATH/src/github.com/tendermint/slate/source/index.html.md

rm $destination

header="---
title: RPC Reference

language_tabs:
  - shell
  - go

toc_footers:
  - <a href='https://tendermint.com/'>Tendermint</a>
  - <a href='https://github.com/lord/slate'>Documentation Powered by Slate</a>

search: true
---"

echo "$header" > "$destination"

rpc_docs=$(godoc2md -template rpc/core/doc_template.txt github.com/tendermint/tendermint/rpc/core | grep -v -e "pipe.go" -e "routes.go" -e "dev.go" | sed 's$/src/target$https://github.com/tendermint/tendermint/tree/master/rpc/core$')

echo "$rpc_docs" >> "$destination"
