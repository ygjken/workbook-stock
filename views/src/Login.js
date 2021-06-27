function Login() {
    return (
      <div className="Login">
        <header className="Login-header">
          <h1>Input login info</h1>
        </header>
        <body>
          <form>
            <fieldset>
                <legend>Legend</legend>
                <div class="form-group row">
                  <label for="staticEmail" class="col-sm-2 col-form-label">Email</label>
                  <div class="col-sm-10">
                    <input type="text" readonly="" class="form-control-plaintext" id="staticEmail" value="email@example.com"></input>
                </div>
                </div>
                <div class="form-group">
                  <label for="exampleInputEmail1" class="form-label mt-4">Email address</label>
                  <input type="email" class="form-control" id="exampleInputEmail1" aria-describedby="emailHelp" placeholder="Enter email"></input>
                  <small id="emailHelp" class="form-text text-muted">We'll never share your email with anyone else.</small>
                </div>
            </fieldset>
          </form>
        </body>
      </div>
    );
  }

  export default Login;