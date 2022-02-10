package template

const HelpMain = `
Here are the available commands:

historical				Gets historical data based on flags
watch							Listen to real time data from all contracts
clean							Clears the local cache

Here are the available flags for the historical command:

--rats						Gets token info, owner, and meta for rats
--closet-pieces		Gets token info and meta for closet pieces
--closet-tokens		Gets owner info on closet tokens
--sync						Syncs data to caching service
--ignore-cache		Ignores local caching files
`
