# Eol-checker

[![made-with-Go](https://img.shields.io/badge/Made%20with-Go-1f425f.svg?style=for-the-badge)](http://golang.org)
[![GitHub go.mod Go version of a Go module](https://img.shields.io/github/go-mod/go-version/gomods/athens.svg?style=for-the-badge)](https://github.com/cheunn-panaa/eol-checker)
[![Go Report Card](https://goreportcard.com/badge/github.com/cheunn-panaa/eol-checker?style=for-the-badge)](https://goreportcard.com/report/github.com/cheunn-panaa/eol-checker)
[![GitHub pull-requests](https://img.shields.io/github/issues-pr/cheunn-panaa/eol-checker?style=for-the-badge)](https://GitHub.com/Cheunn-Panaa/eol-checker/pull/)
[![GitHub license](https://img.shields.io/github/license/cheunn-panaa/eol-checker?style=for-the-badge)](https://github.com/Cheunn-Panaa/eol-checker/blob/main/LICENSE)

<!-- ABOUT THE PROJECT -->
## 🧙‍♂️ About The Project 🧙‍♂️

[![Product Name Screen Shot](
https://static.wikia.nocookie.net/injusticegodsamongus/images/2/2c/Detective_Chimp_Injustice_Y3.jpg/revision/latest?cb=20141006022951)](https://google.com/)

End of life is not fun, please check your apps once in a while.

Here's why:
* Legacy is no bueno
* Upgrading takes time
* Non-tech individuals are dumbdumbs and will probably not understand why it costs a fair amount of money.

Of course, you are free to do whatever you want.

<p align="right">(<a href="#top">back to top</a>)</p>


<!-- GETTING STARTED -->
## Getting Started 

This is an example of how you may give instructions on setting up your project locally.
To get a local copy up and running follow these simple example steps.

I need to do a proper makefile.... yikes

Anyway ...
- [Installation](docs/INSTALLATION.md)
- [Configuration](docs/CONFIGURATION.md)

<!-- USAGE EXAMPLES -->
## Usage

The usage of this cli is pretty straight forward.
Once you have setup the cli properly, **checkout [Getting Start](#getting-started) if you have missed it**, you can simply run


Simply run ```eol-checker```
```eol-checker -h``` will show usage and list of command

```bash
EOL is a CLI library for project version management. 
 This application is a tool to check wether or not your application version is out of date. 
For now it checks on endoflife.date endpoints but is subject to change in the future

Usage:
  eol-cli [flags]

Flags:
  -c, --config string     Configuration file to use (default "eol-cli.yaml")
  -d, --disable-message   Disables the notifications on all plugins
  -h, --help              help for eol-cli
  -v, --version           Display the current version of this CLI
```

_For more examples, please refer to the [Documentation](https://example.com)_

<p align="right">(<a href="#top">back to top</a>)</p>



<!-- ROADMAP -->
## Roadmap

- [x] Add A Readme template
- [x] Add back to top links
- [x] Make it runnable as a cli ?
    - [ ] Make it runnable project wise rather than global.
    - [ ] Is global conf still necessary 

- [ ] Better docs (but i wont)
- [ ] Proper Error Handling
- [ ] More version checking
    - [ ] framework specific ?
    - [ ] library ?
    - [ ] Dockerfile ?
    - [ ] ???
- [ ] More "Plugins" for alerting
    - [ ] Mailing
    - [ ] Telegram ?
    - [ ] Discord ?
    - [ ] ????
- [ ] Better roadmap ?

See the [open issues](https://github.com/Cheunn-Panaa/eol-checker/issues) for a full list of proposed features (and known issues).

<p align="right">(<a href="#top">back to top</a>)</p>



<!-- CONTRIBUTING -->
## Contributing

Contributions are what make the open source community such an amazing place to learn, inspire, and create. Any contributions you make are **greatly appreciated**.

Only contribute if you are very good in golang so you can fix my rookie mistakes.

<p align="right">(<a href="#top">back to top</a>)</p>



<!-- LICENSE -->
## License

Distributed under the AGPL-3.0 License. See `LICENSE` for more information.
idk what License means ?????

<p align="right">(<a href="#top">back to top</a>)</p>