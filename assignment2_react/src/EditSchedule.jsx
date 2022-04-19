import React, { useState } from "react";
import { toast } from "react-toastify";

function EditSchedule(props) {
  const [scheduleEditMovieText, setScheduleEditMovieText] = useState(
    props.movie
  );
  const [scheduleEditVenueText, setsSheduleEditVenueText] = useState(
    props.venue
  );
  const [scheduleEditDateText, setScheduleEditDateText] = useState(
    props.datehour
  );

  var element = document.getElementById(props.venue + props.movie)

  const onClick = async (e) => {
    console.log(
      scheduleEditMovieText,
      scheduleEditVenueText,
      scheduleEditDateText
    );
    // const prefix = "http://localhost:8080";
    const prefix = "https://goschoolassignment2.onrender.com";
    const url = `${prefix}/singledatehour/${props.datehour}?movie=${props.movie}&venue=${props.venue}`;
    const result = await fetch(url, {
      method: "PUT",
      headers: {
        "Content-Type": "application/json",
        Origin: "http://localhost:3000",
      },
      body: JSON.stringify({
        newdateHour: parseInt(scheduleEditDateText),
        newvenue: scheduleEditVenueText,
        newmovie: scheduleEditMovieText,
      }),
    });
    const tmp = await result.json();
    // window.location.reload();
    if (!result.ok) {
      toast.error(tmp.error);
      return;
    }
    if (tmp.status === "ok") {
      setScheduleEditMovieText("");
      setsSheduleEditVenueText("");
      window.location.reload();
    }
    e.preventDefault();
  };

  const onChangeInputScheduleMovie = (e) => {
    setScheduleEditMovieText(e.target.value);
  };

  const onChangeInputScheduleVenue = (e) => {
    setsSheduleEditVenueText(e.target.value);
  };

  const onChangeInputScheduleDateHour = (e) => {
    setScheduleEditDateText(e.target.value);
  };

  return (
    <>
      <button
        type="button"
        class="btn btn-secondary btn-sm"
        data-bs-toggle="modal"
        data-bs-target={"#" + props.venue + props.movie}
      >
        Edit
      </button>
      <div class="modal fade" id={props.venue + props.movie} tabIndex="-1">
        <div class="modal-dialog">
          <div class="modal-content">
            <div class="modal-header">
              <h5 class="modal-title" id="editScheduleLabel">
                Edit Schedule
              </h5>
              <button
                type="button"
                class="btn-close"
                data-bs-dismiss="modal"
                aria-label="Close"
              ></button>
            </div>
            <form>
              <div class="mb-3">
                <label for="movie" class="form-label">
                  DateHour
                </label>
                <input
                  type="text"
                  class="form-control"
                  id="movie"
                  value={scheduleEditDateText}
                  onChange={onChangeInputScheduleDateHour}
                />
              </div>
              <div class="mb-3">
                <label for="movie" class="form-label">
                  Edit movie name
                </label>
                <div class="form-text">New movie must be added first</div>
                <input
                  type="text"
                  class="form-control"
                  id="movie"
                  value={scheduleEditMovieText}
                  onChange={onChangeInputScheduleMovie}
                />
              </div>
              <div class="mb-3">
                <label for="venue" class="form-label">
                  Edit venue name
                </label>
                <div class="form-text">New venue must be added first</div>
                <input
                  type="text"
                  class="form-control"
                  id="movie"
                  value={scheduleEditVenueText}
                  onChange={onChangeInputScheduleVenue}
                />
              </div>
            </form>

            <div class="modal-footer">
              <button
                type="button"
                class="btn btn-secondary"
                data-bs-dismiss="modal"
              >
                Close
              </button>
              <button
                type="button"
                class="btn btn-primary btn-sm"
                onClick={onClick}
              >
                Save changes
              </button>
            </div>
          </div>
        </div>
      </div>
    </>
  );
}

export default EditSchedule;
