# calcy

A study in polymorphic deployment of a package across multiple UIs.

1. `calcli` - bare-bones CLI (no flags, addition-only)
2. `calcli2` - many supported operations (via flags package) 
3. `calcsv` - process a 3-column csv file
4. `calhttp` - http server w/ routes for each operation
5. `calcsv2http` - client reads/writes csv file, getting answers from `calhttp` server