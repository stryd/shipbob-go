openapi-generator generate -i shipbob_open_api.json -g go --additional-properties=packageName=shipbob,packageVersion=1.2,disallowAdditionalPropertiesIfNotPresent=false --git-repo-id shipbob-go --git-user-id stryd --global-property apis,apiTests=false,,models,modelTests=false,supportingFiles