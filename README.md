<p align="center">
    <img alt="Apes Logo" src="https://raw.githubusercontent.com/Clivern/Apes/master/assets/img/gopher.png?v=0.1.1" width="150" />
    <h3 align="center">Apes</h3>
    <p align="center">Chaos and Resiliency Testing Service.</p>
    <p align="center">
        <a href="https://travis-ci.com/Clivern/Apes"><img src="https://travis-ci.com/Clivern/Apes.svg?branch=master"></a>
        <a href="https://github.com/Clivern/Apes/releases"><img src="https://img.shields.io/badge/Version-0.1.1-red.svg"></a>
        <a href="https://goreportcard.com/report/github.com/Clivern/Apes"><img src="https://goreportcard.com/badge/github.com/clivern/Apes?v=0.1.1"></a>
        <a href="https://github.com/Clivern/Apes/blob/master/LICENSE"><img src="https://img.shields.io/badge/LICENSE-MIT-orange.svg"></a>
    </p>
</p>

Apes can work as a proxy for one of your upstream API services to simulate high latencies and failures to make sure your services have the capability to withstand and recover from failures.

## Documentation

### Usage

Get [the latest binary.](https://github.com/Clivern/Apes/releases)

```zsh
$ curl -sL https://github.com/Clivern/Apes/releases/download/x.x.x/Apes_x.x.x_OS_x86_64.tar.gz | tar xz
```

Run Apes Chaos Reverse Proxy.

```zsh
$ ./Apes --port=8080 --upstream=https://httpbin.org --failRate=10% --latency=0s
```

Check the release.

```zsh
$ ./Apes --get=release
```

Test it.

```zsh
$ curl http://127.0.0.1:8080/ip
```


## Versioning

For transparency into our release cycle and in striving to maintain backward compatibility, Apes is maintained under the [Semantic Versioning guidelines](https://semver.org/) and release process is predictable and business-friendly.

See the [Releases section of our GitHub project](https://github.com/clivern/apes/releases) for changelogs for each release version of Apes. It contains summaries of the most noteworthy changes made in each release.


## Bug tracker

If you have any suggestions, bug reports, or annoyances please report them to our issue tracker at https://github.com/clivern/apes/issues


## Security Issues

If you discover a security vulnerability within Apes, please send an email to [hello@clivern.com](mailto:hello@clivern.com)


## Contributing

We are an open source, community-driven project so please feel free to join us. see the [contributing guidelines](CONTRIBUTING.md) for more details.


## License

© 2020, clivern. Released under [MIT License](https://opensource.org/licenses/mit-license.php).

**Apes** is authored and maintained by [@clivern](http://github.com/clivern).
