# Lazy Hangout
## Introduction
**Lazy Hangout** is a simple utility that aims to automate the generation of 
a Google Hangout link. Often, I tend to converse with my friends over Google
Hangout and it feels quite cumbersome to navigate through my web browser in
order to start a Hangout chat.

## Get Started
Please follow the instructions step-by-step as it will aid the process of setting 
up this utility.

### Pre-requisites
- Go

### Service Account
1. Create a [service account]((https://cloud.google.com/iam/docs/service-accounts)) 
for **Google Calendar API**
2. By now, you should get a `.json` file with your service account's credentials
3. Place your credential file anywhere that *feels* secure to you

### Environment Variable
The absolute file path in `HANGOUT_CREDENTIALS_PATH` is used to locate your credential 
file. 

```bash
export HANGOUT_CREDENTIALS_PATH="<absolute-path-to-your-credential-file>"
```

### Build
The last step is to clone this repository and build the project.

```bash
git clone https://github.com/LeonLyx/lazy-hangout.git
cd lazy-hangout
make build
```

Once done, you should be able to generate a hangout link by executing `hangout-generator` 
on your favourite terminal.



