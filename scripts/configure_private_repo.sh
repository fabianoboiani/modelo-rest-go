#!/bin/bash

git config \
  --global \
  url."https://lojasrenner:35w7kwespmgnsiqvgjusa2s4up7pqaw74qy2h6peyuwjauguf6ba@dev.azure.com/lojasrenner/Data%20Engineering/_git/datalab-commons".insteadOf \
  "https://dev.azure.com/lojasrenner/datalab-commons"

go env -w GOPRIVATE=dev.azure.com/lojasrenner/*