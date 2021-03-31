curl -X POST \
      -H 'Circle-Token: a7ba232028b86cff5396141a7dfa7f25bf451eba' \
      -H 'Content-Type: application/json' \
      -H 'Accept: application/json' \
      -d "{\"branch\":\"INS-12888-Enable-CircleCI-Build-Test-On-Commit\", \"parameters\":{\"is-pr\":true}}" \
      https://circleci.com/api/v2/project/github/EngineersBox/terraform-provider-instaclustr/pipeline