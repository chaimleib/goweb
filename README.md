# goweb

This is a hello world for a go-based web server. To set things up:

1. Download the latest version of repoactions from [https://github.com/chaimleib/repoactions/releases/latest]()
2. Unpack it and from the terminal, run this inside the repoactions folder:

  ```bash
  ./configure
  make install
  source ~/.profile
  ```

3. Clone the goweb repo and cd into it.

  ```bash
  git clone git@github.com:chaimleib/goweb.git
  cd goweb
  ```

4. Follow the instructions repoactions gives to add goweb to the repoactions whitelist. This will set up your GOPATH and any other needed environment variables for you whenever you cd into the goweb repo.

5. Run the tests:

  ```bash
  make test
  ```

6. Compile the server. This compiles the code from `src/server/cli/server/main.go` and puts the binary in `bin/server`.

  ```bash
  make
  ```

7. Run the server. Thanks to repoactions, the `bin` folder is now in your PATH, so you can just type this:

  ```bash
  server
  ```

8. Navigate to [http://localhost:8080/]() to see the hello world!
