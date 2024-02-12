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
npx golagen [config_file_name (default: "configs/golagen.yml")]
```

> yarn
```bash
yarn add --dev@lucas-cox/golagen
yarn run golagen [config_file_name (default: "configs/golagen.yml")]
```


Make sure that a `configs` directory exists, with a valid golagen configuration file.

### Global Installation
Use the following command to install **golagen** globally and run it :
> npm
```bash
npx golagen [config_file_name (default: "golagen.yml")]
```

> yarn
```bash
yarn global add @lucas-cox/golagen
golagen [config_file_name (default: "golagen.yml")]
```

Make sure that a `.golagen` directory exists in the root of your project directory, with a valid `main.yml` file.


## In development

- Generate a go lambda monorepo using an **aws-sam** based architecture, generating the [SAM template file](https://docs.aws.amazon.com/codedeploy/latest/userguide/tutorial-lambda-sam-create-lambda-function.html), wrapped in a Makefile
- Generate a Makefile calling all the lambda ones and regenerating them if needed
- Add custom rules to the generated Makefile for each entry and for the global one
- Deploy particular lambdas, using aws-sam if possible else aws-cli