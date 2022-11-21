### Small Go application that groups lines of an input file by their first letter and exports the result to an output file

#### How to run

- Download and install Go from their [website](https://go.dev/dl)
- Download the files from this repository on your device
- Navigate to the _"main"_ folder
- To run the app, open the command line in the folder and type **"go run main.go"**. The result will be exported in _"output.csv"_
- To run the tests, type **"go test"**
- If you wish an other input, change the contents inside _"input.csv"_. Be careful to respect the [format](#format) below or the app will crash

#### Format

- The first line should contain lowercase letters
- The next lines should start with with an uppercase letter
- Duplicated lines can exist, but they will be ignored

#### Example

== _input.csv_ ==<br>

full_name, email, location<br>
Anita, anita<span>@</span>email.com, California<br>
Aron, aron.bla<span>@</span>email.com, California<br>
Cosmin, kox<span>@</span>bla.com, Giurgiu<br>
Aron, aron.bla<span>@</span>email.com, California<br>

<br>

== _output.csv_ ==<br>

A:<br>
Anita, anita<span>@</span>email.com, California<br>
Aron, aron.bla<span>@</span>email.com, California<br>

C:<br>
Cosmin, kox<span>@</span>bla.com, Giurgiu<br>
