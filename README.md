# Prom-Go

A go helper library that contains the alert struct for prometheus alerts

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes. See deployment for notes on how to deploy the project on a live system.

### Prerequisites

- golang

```
yum/apt-get/brew install go
```

### Installing

Import the library into your go application

```
import (
	"github.com/HealthPartnersOSS/prom-go"
)
```

Use the struct to create a new AlertGroup, or parse a message into the struct

```
alertGroup := prometheus.AlertGroup{
    Version:        "1.2.3",
	GroupKey:       "my-group",
	Status:         "resolved",
	Receiver:       "my-receiver",
	ExternalURL:    "http://example.com",
}

b := []byte(j)

var m prometheus.AlertGroup
err := json.Unmarshal(b, &m)
```

## Built With

* [Golang](https://golang.org/)

## Contributing

Please read [CONTRIBUTING.md](https://gist.github.com/PurpleBooth/b24679402957c63ec426) for details on our code of conduct, and the process for submitting pull requests to us.

## Authors

* **Peter Kreidermacher** - *Initial work* - [HealthPartners](https://github.com/healthpartnersoss)

See also the list of [contributors](https://github.com/healthpartnersoss/prom-go/contributors) who participated in this project.

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details
