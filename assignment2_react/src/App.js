import Calendar from "react-calendar";
import React, { useState, useEffect } from "react";
import VenueEdit from "./VenueEdit";
import MovieEdit from "./MovieEdit";
import NewSchedule from "./NewSchedule";
import "react-calendar/dist/Calendar.css";
import { ToastContainer, toast } from "react-toastify";
import "react-toastify/dist/ReactToastify.css";
import EditSchedule from "./EditSchedule";

function compare(a, b) {
  if (a.DateHour < b.DateHour) {
    return -1;
  }
  if (a.DateHour > b.DateHour) {
    return 1;
  }
  return 0;
}

function App() {
  const [value, onChange] = useState(new Date());
  const [data, setData] = useState("");
  const [deleteItem, setDeleteItem] = useState("");
  const [venues, setVenues] = useState("");
  const [movies, setMovies] = useState("");
  const [searchByVenue, setSearchByVeneue] = useState("");
  const [searchByMovies, setSearchByMovies] = useState("");

  var year = value.getFullYear();
  var month = value.getMonth() + 1;
  if (month.toString().length === 1) {
    month = "0" + month;
  }
  var day = value.getDate();
  var startDate = year.toString() + month.toString() + day.toString() + "00";
  var endDate = year.toString() + month.toString() + day.toString() + "23";
  // console.log(startDate, endDate);

  // const prefix = "http://localhost:8080";
  const prefix = "https://goschoolassignment2.onrender.com";

  useEffect(() => {
    const getAPI = async () => {
      let url;
      url = `${prefix}/rangedatehour/?start=${startDate}&end=${endDate}`;

      const result = await fetch(url);

      if (!result.ok) {
        throw new Error("Something went wrong");
      }

      const tmp = await result.json();
      //   console.log(word);
      // console.log(data);
      setData(tmp);
    };
    getAPI();
  }, [startDate, endDate, deleteItem]);

  useEffect(() => {
    const getVenues = async () => {
      let url;
      url = `${prefix}/venues`;

      const result = await fetch(url);

      if (!result.ok) {
        throw new Error("Something went wrong");
      }

      const tmp = await result.json();
      setVenues(tmp);
    };
    getVenues();
  }, []);

  useEffect(() => {
    const getMovies = async () => {
      let url;
      url = `${prefix}/movies`;

      const result = await fetch(url);

      if (!result.ok) {
        throw new Error("Something went wrong");
      }

      const tmp = await result.json();
      //   console.log(word);
      // console.log(data);
      setMovies(tmp);
    };
    getMovies();
  }, []);

  const onClick = async (e) => {
    // console.log(e.target.id);
    if (window.confirm("Are you sure to delete this event?") === true) {
      const dh = e.target.id.split("_")[0];
      const ve = e.target.id.split("_")[1];
      const mo = e.target.id.split("_")[2];
      setDeleteItem({ dh, ve, mo });
      const url = `${prefix}/singledatehour/${dh}?venue=${ve}&movie=${mo}`;
      const result = await fetch(url, { method: "DELETE" });

      if (!result.ok) {
        throw new Error("Something went wrong");
      }

      e.preventDefault();
    }
  };

  const onClickBalance = async (e) => {
    if (window.confirm("Are you sure to balance the current tree?") === true) {

      const url = `${prefix}/balance`;
      const result = await fetch(url);
      const tmp = await result.json();

      // console.log(tmp)
      // console.log(result)

      if (!result.ok) {
        throw new Error("Something went wrong");
      } else {
        toast.success("Balance Successfully!" );
      }

      e.preventDefault();
    }
  };

  const onClickVeneue = async (e) => {
    // console.log(e.target.id);
    const ve = e.target.id;
    // console.log(id)
    const url = `${prefix}/rangedatehour/?venue=${ve}`;
    const result = await fetch(url);
    if (!result.ok) {
      throw new Error("Something went wrong");
    }
    const tmp = await result.json();
    setSearchByVeneue(tmp);
    e.preventDefault();
  };

  const onClickMovie = async (e) => {
    // console.log(e.target.id);
    const ve = e.target.id;
    // console.log(id)
    const url = `${prefix}/rangedatehour/?movie=${ve}`;
    const result = await fetch(url);
    if (!result.ok) {
      throw new Error("Something went wrong");
    }
    const tmp = await result.json();
    setSearchByMovies(tmp);
    e.preventDefault();
  };

  return (
    <>
      <nav className="navbar navbar-expand-lg navbar-light bg-light">
        <div className="container-fluid">
          <a className="navbar-brand" href="#">
            GoSchool Assignment2
          </a>
          <button
            className="navbar-toggler"
            type="button"
            data-bs-toggle="collapse"
            data-bs-target="#navbarNavAltMarkup"
            aria-controls="navbarNavAltMarkup"
            aria-expanded="false"
            aria-label="Toggle navigation"
          >
            <span className="navbar-toggler-icon"></span>
          </button>
          <div className="collapse navbar-collapse" id="navbarNavAltMarkup">
            <div className="navbar-nav"></div>
          </div>
        </div>
        <form class="form-inline">
          <button
            class="btn btn-sm btn-outline-secondary"
            type="button"
            onClick={onClickBalance}
          >
            Click me to Balance Tree!
          </button>
        </form>
      </nav>

      <div className="container">
        <h1>Movies Schedules</h1>
        <hr></hr>
        <br></br>
        <div className="row">
          <div className="col-sm-3">
            <Calendar onChange={onChange} value={value} />
          </div>

          {/* <div className="flex-item"> */}
          <div className="col-sm-3">
            {!data.message && <div>No Movie for the day!</div>}
            {data.message &&
              data.message.map((item) => (
                <div
                  className="card"
                  style={{ width: "18rem" }}
                  key={item.DateHour + "_" + item.Venue + "_" + item.Movie}
                >
                  <div className="card-body">
                    <div>
                      Date: {item.DateHour.toString().slice(0, 8)}, Time:{" "}
                      {item.DateHour.toString().slice(8, 10)}:00
                    </div>
                    <div>Venue: {item.Venue}</div>
                    <div>Movie Name: {item.Movie}</div>
                    <button
                      type="button"
                      className="btn btn-primary btn-sm"
                      id={item.DateHour + "_" + item.Venue + "_" + item.Movie}
                      onClick={onClick}
                    >
                      Delete
                    </button>
                    <EditSchedule
                      datehour={item.DateHour}
                      venue={item.Venue}
                      movie={item.Movie}
                    ></EditSchedule>
                  </div>
                </div>
              ))}
          </div>
          <div className="col-sm-3">
            <NewSchedule></NewSchedule>
          </div>
        </div>
        <br></br>
        <br></br>
        <br></br>
        <br></br>
        <div className="row">
          <div className="col-sm-2">
            <h5>Venue Lists:</h5>
            <VenueEdit></VenueEdit>
            <div className="card small">
              <div className="card-body small">
                {!venues.message && <div>No Venues!</div>}
                {venues.message &&
                  venues.message.map((item) => (
                    <div key={item} class="searchbutton">
                      {item}{" "}
                      <button
                        type="button"
                        className="btn btn-light right btn-sm"
                        id={item}
                        onClick={onClickVeneue}
                      >
                        List all
                      </button>
                    </div>
                  ))}
              </div>
            </div>
          </div>
          <div className="col-sm-2">
            <h5>Movies Lists:</h5>
            <MovieEdit></MovieEdit>
            <div className="card small">
              <div className="card-body small">
                {!movies.message && <div>No Movies!</div>}
                {movies.message &&
                  movies.message.map((item) => (
                    <div key={item} class="searchbutton">
                      {item}{" "}
                      <button
                        type="button"
                        className="btn btn-light right btn-sm"
                        id={item}
                        onClick={onClickMovie}
                      >
                        List all
                      </button>
                    </div>
                  ))}
              </div>
            </div>
          </div>
          <div className="col-sm-4">
            <div className="card">
              <div className="card-body">
                <h5>Search By Venue:</h5>
                {!searchByVenue.message && <div>No Search Result</div>}
                {searchByVenue.message &&
                  searchByVenue.message.sort(compare).map((item) => (
                    <div
                      key={item.DateHour + "_" + item.Venue + "_" + item.Movie}
                      className="fixed2"
                    >
                      <div>
                        {item.DateHour}; {item.Venue}; {item.Movie}
                      </div>
                    </div>
                  ))}
              </div>
            </div>
          </div>
          <div className="col-sm-4">
            <div className="card">
              <div className="card-body">
                <h5>Search By Movies:</h5>
                {!searchByMovies.message && <div>No Search Result</div>}
                {searchByMovies.message &&
                  searchByMovies.message.sort(compare).map((item) => (
                    <div
                      key={item.DateHour + "_" + item.Venue + "_" + item.Movie}
                    >
                      <div>
                        {item.DateHour}; {item.Venue}; {item.Movie}
                      </div>
                    </div>
                  ))}
              </div>
            </div>
          </div>
        </div>
      </div>
      <ToastContainer />
    </>
  );
}

export default App;
