# memorize
Command-line flashcards program.

> memorize is currently in development. Expect its behavior to change. Contributions are welcome.

## Usage

Flashcards are a set of key-value pairs stored in a two/three column [TSV file][1]. memorize accepts the flashcards file path as an argument, and reads the file line by line, prompting the user with the first term in a line, and asking for the corresponding second term.

If you have a file called `vocab.tab` that looks like this:

```
aimer bien	to like	eh-meh byah
apporter	to bring
coûter	to cost
à emporter	take out	ah om-por-teh
faire des courses	to run errands
```

A memorize session should look like this:

```
$ memorize vocab.tab
aimer bien: to like
Correct! [eh-meh byah]
apporter: to bring
Correct!
coûter: to cost
Correct!
à emporter: take away
take out [ah om-por-teh]
faire des courses: to go grocery shopping
to run errands
$ 
```

Flashcards may have a "third side", which is displayed together with the right/wrong message (see above for an example).

The current implementation doesn't permit fast memorization of terms. Future work will require implementation of the [Leitner flashcard system][2] for effective information acquisition.


[1]: https://en.wikipedia.org/wiki/Tab-separated_values
[2]: https://en.wikipedia.org/wiki/Leitner_system