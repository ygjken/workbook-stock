//import logo from './logo.svg';
import './App.css';
import './nav.css';

function App() {
  return (
    <div className="Welcome">
      <header className="Welcome-header">
        <h1>Welcome Our Webpage!!</h1>
      </header>
      <body>
        <nav class="navbar navbar-expand-lg navbar-dark bg-primary">
          <div class="container-fluid">
            <a class="navbar-brand" href="#">Navbar</a>
            <button class="navbar-toggler" type="button" data-bs-toggle="collapse" data-bs-target="#navbarColor01" aria-controls="navbarColor01" aria-expanded="false" aria-label="Toggle navigation">
              <span class="navbar-toggler-icon"></span>
            </button>

            <div class="collapse navbar-collapse" id="navbarColor01">
              <ul class="navbar-nav me-auto">
                <li class="nav-item">
                  <a class="nav-link active" href="#">Home
                    <span class="visually-hidden">(current)</span>
                  </a>
                </li>
                <li class="nav-item">
                  <a class="nav-link" href="#">LogIn</a>
                </li>
                <li class="nav-item">
                  <a class="nav-link" href="#">Coming soon1</a>
                </li>
                <li class="nav-item">
                  <a class="nav-link" href="#">Coming soon2</a>
                </li>
                <li class="nav-item dropdown">
                  <a class="nav-link dropdown-toggle" data-bs-toggle="dropdown" href="#" role="button" aria-haspopup="true" aria-expanded="false">Dropdown</a>
                  <div class="dropdown-menu">
                    <a class="dropdown-item" href="#">Action</a>
                    <a class="dropdown-item" href="#">Another action</a>
                    <a class="dropdown-item" href="#">Something else here</a>
                    <div class="dropdown-divider"></div>
                    <a class="dropdown-item" href="#">Separated link</a>
                  </div>
                </li>
              </ul>
              <form class="d-flex">
                <input class="form-control me-sm-2" type="text" placeholder="Search"></input>
                <button class="btn btn-secondary my-2 my-sm-0" type="submit">Search</button>
              </form>
            </div>
          </div>
        </nav>
      </body>
    </div>
  );
}

export default App;
