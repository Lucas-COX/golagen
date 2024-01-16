<h1 align="center" style="border-bottom: none">Golagen</h1>
<h3 align="center">Project generator based on Terraform and AWS lambda</h3>

<p align="center">
    <a href="https://github.com/semantic-release/semantic-release">
        <img alt="Semantic Release" src="https://img.shields.io/badge/semantic--release-angular-e10079?logo=semantic-release">
    </a>
    <a href="https://github.com/Lucas-COX/golagen/releases">
        <img alt="Last release" src="https://img.shields.io/github/v/release/Lucas-COX/golagen">
    </a>
    <a href="https://github.com/Lucas-COX/golagen/releases">
        <img alt="Release date" src="https://img.shields.io/github/release-date/Lucas-COX/golagen">
    </a>
</p>

The goal of this project is to easily generate different types of projects (lambda monorepos, react and next apps, etc.) with the idea of being highly configurable.

Please note that this is not functional and being developped for the moment.


## Requirements
In order to use golagen, you need the following packages :
- npm >= 5.2.0
- go >= 1.21.0


## Installation

### Local Installation

Use the following commands to install **golagen** directly in your project's directory and run it :
> npm
```bash
npm install --save-dev @lucas-cox/golagen
npx golagen
```

> yarn
```bash
yarn add --dev@lucas-cox/golagen
yarn run golagen
```


Make sure that a `.golagen` directory exists, with a valid `main.yml` file.

### Global Installation
Use the following command to install **golagen** globally and run it :
> npm
```bash
npx golagen [path_to_project]
```

> yarn
```bash
yarn global add @lucas-cox/golagen
golagen [path_to_project]
```

Make sure that a `.golagen` directory exists in the root of your project directory, with a valid `main.yml` file.


## In development

- Generate a lambda monorepo using a terraform friendly architecture, giving the sample terraform files, wrapped in a Makefile
- Generate a Makefile calling all the projects ones
- Parse the configuration file before generating to avoid catching exceptions during generation
- Make some configuration parameters mandatory
- Add custom rules to the generated Makefile for each entry and for the global one
- Let user choose if script creates a Makefile rule or not