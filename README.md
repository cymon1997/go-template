# go-template

Minimal project structure for building [Go](https://golang.org) application.

## Features

- Clean `main.go` initialization
- Provider based dependency injection (using singleton)
- Config files replaceable with ENV variables 
- Core layer separation: inbound, outbound, usecase 

## Installation

```bash
go get -u github.com/cymon1997/go-template
```

## Contribute

### Development

Checkout from latest `main` branch
```bash
git checkout main 
git pull origin main 
git checkout -b <your_branch>
```
Hint: please take a look at [Branch Convention](#branch-convention)

If you add other dependencies, run:
```bash
make update-dep 
```

Before raise a Pull Request, please make sure you already suffice the tests of your code.

```bash
make test
```

### Branch Convention

Format:
> [prefix]_[feature_name]

Prefix:
- `f_` for new feature implementation
- `i_` for adding code improvement
- `b_` for fixing bug

Examples:
- f_i
- i_setup_http_transport
- b_fix_intermittent_http

### Issue / Feature Request

Please raise an issue and explains the issue / feature that you want to be supported.
Give detail explanation about the issue / feature.

## Contact

If you have anything to ask / discuss, please contact me below, thanks!   
Aji Imawan Omi  
GitHub: cymon1997

## License

GNU GENERAL PUBLIC LICENSE - Aji Imawan Omi