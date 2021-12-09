# Organisation API

This is a client API for the `accounts` resource under Organisation defined on the [Form3 documentation](http://api-docs.form3.tech/api.html#organisation-accounts).

Couple of assumptions were made:

- It was mentioned that the tests should run against the mock Account API, it was assumed that this meant that there should be integration tests.
- Not every test runs against the mock Account API, there are normal unit tests, since it also mentioned that it should be tested as in a prod environment.

Created by Cristiano Guilherme de Souza Silva, email: `criguiss@uol.com.br`

I've used Go on a professional manner, however, not to a great extent. 

Thanks to Form3 for the opportunity :)

## Structure

```bash
organisation-api
    ├───.idea
    └───scripts
       └───db
```

