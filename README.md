## Assessment Task

### Task 1 Complexity & Estimation

This task is basically dependent on some utility functions like splitting text letteral into array of unsigned numbers and strings.

The estimated time is 5 hours as we need to write unit test cases also. This task basically all about implementation and making strong logic and writing optimized, clear and non-redundant code.

This was the main reason that this took more time as compare to the unit testing.

### Task 2 Difficulty & Estimation

This task is basically related to unit testing. I have just done the unit testing portion and skipped the first portion as the requirement is a bit confusing for me. The word **random correct strings**, I don't know what do you mean by this.

### Installation

Clone the repository from the link `git@github.com:hussain2y2/lucky-jobs.git` if we are using `SSH` as the https support is abondend after August 8, 2020. So either we nedd to use `SSH` or Github CLI: `gh repo clone hussain2y2/lucky-jobs`.

## Running

After cloning is successful, now its time to build and run it.

1. Use the command `go build` or `go build -o <your desired name>` after changing the directory into the local repository to build binary.
2. After building the binary there will be a file generated with name `lucky-jobs`. Run the command `./lucky-jobs` or the `./<your desired name>` and we will see the out based on the input text we changed in the main.go file. The current one is `23-ab-48-caba-56-haha`.
3. Run the command `go test -v` to check if all the test cases written are working fine.
