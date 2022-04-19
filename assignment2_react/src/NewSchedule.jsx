import React, { useState } from "react";
import { toast } from "react-toastify";

function NewSchedule() {
  const [newDateHour, setNewDateHour] = useState("");
  const [newVenue, setNewVenue] = useState("");
  const [newMovie, setNewMovie] = useState("");

  const onChangeDateHour = (e) => {
    setNewDateHour(e.target.value);
  };

  const onChangeVenue = (e) => {
    setNewVenue(e.target.value);
  };

  const onChangeMatch = (e) => {
    setNewMovie(e.target.value);
  };

  const onClickAddSchedule = async (e) => {
    e.preventDefault();
    if (newDateHour === "" || newVenue === "" || newMovie === "") {
        toast.error("Please fill all the fields");
        return;
    }
    // const prefix = "http://localhost:8080";
    const prefix = "https://goschoolassignment2.onrender.com";
    const url = `${prefix}/singledatehour`;
    const result = await fetch(url, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
        Origin: "http://localhost:3000",
      },
      body: JSON.stringify({
        dateHour: parseInt(newDateHour),
        venue: newVenue,
        movie: newMovie,
      }),
    });
    const tmp = await result.json();
    if (!result.ok) {
      toast.error(tmp.error);
      setNewDateHour("");
      setNewVenue("");
      setNewMovie("");
      return;
    }
    if (tmp.status === "ok") {
      // setVenues(tmp.data);
      setNewDateHour("");
      setNewVenue("");
      setNewMovie("");
      toast.success("Add new schedule successfully!");
    }
    
  };

  return (
    <>
      <form>
        <div class="mb-3">
          <label for="datehour" class="form-label">
            Date & Hour
          </label>
          <input
            type="text"
            class="form-control"
            id="datehour"
            value={newDateHour}
            onChange={onChangeDateHour}
            required
          />
          <div id="datehour" class="form-text">
            Please enter the datehour of the new movie schedule, example
            "2022041518"
          </div>
        </div>
        <div class="mb-3">
          <label for="venue" class="form-label">
            Venue
          </label>
          <input
            type="text"
            class="form-control"
            id="venue"
            value={newVenue}
            onChange={onChangeVenue}
            required
          />
          <div id="venue" class="form-text">
            Venue must exist first
          </div>
        </div>
        <div class="mb-3">
          <label for="movie" class="form-label">
            Movie
          </label>
          <input
            type="text"
            class="form-control"
            id="movie"
            value={newMovie}
            onChange={onChangeMatch}
            required
          />
          <div id="movie" class="form-text">
            Movie must exist first
          </div>
        </div>
        <button
          type="submit"
          class="btn btn-primary btn-sm"
          onClick={onClickAddSchedule}
        >
          Add New Schedule
        </button>
      </form>
    </>
  );
}

export default NewSchedule;
