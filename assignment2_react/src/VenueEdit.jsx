import React, { useState } from "react";
import { toast } from "react-toastify";

function VenueEdit() {
  // const prefix = "http://localhost:8080";
  const prefix = "https://goschoolassignment2.onrender.com";
  const [venueEditText, setvenueEditText] = useState("");
  const [venueDeleteText, setVenueDeleteText] = useState("");
  const onChangeInputEdit = (e) => {
    setvenueEditText(e.target.value);
  };

  const onChangeInputDelete = (e) => {
    setVenueDeleteText(e.target.value);
  };

  const onClickAddVenue = async (e) => {
    const url = `${prefix}/venues`;
    const result = await fetch(url, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
        Origin: "http://localhost:3000",
      },
      body: JSON.stringify({
        venue: venueEditText,
      }),
    });
    const tmp = await result.json();
    if (!result.ok) {
      toast.error(tmp.error);
      setvenueEditText("");
      return;
    }
    if (tmp.status === "ok") {
      // setVenues(tmp.data);
      setvenueEditText("");
      window.location.reload();
    }
    e.preventDefault();
  };

  const onClickDelVenue = async (e) => {
    const url = `${prefix}/venues/${venueDeleteText}`;
    const result = await fetch(url, {
      method: "DELETE",
      headers: {
        Origin: "http://localhost:3000",
      },
    });
    const tmp = await result.json();
    if (!result.ok) {
      toast.error(tmp.error);
      setVenueDeleteText("");
      return;
    }
    if (tmp.status === "ok") {
      setVenueDeleteText("");
      window.location.reload();
    }
    e.preventDefault();
  };

  return (
    <>
      <button
        type="button"
        class="btn btn-secondary btn-sm"
        data-bs-toggle="modal"
        data-bs-target="#editVenue"
      >
        Edit Venue
      </button>
      <div class="modal fade" id="editVenue" tabIndex="-1">
        <div class="modal-dialog">
          <div class="modal-content">
            <div class="modal-header">
              <h5 class="modal-title" id="editVenueLabel">
                Edit Venue
              </h5>
              <button
                type="button"
                class="btn-close"
                data-bs-dismiss="modal"
                aria-label="Close"
              ></button>
            </div>
            <div class="modal-body">
              <form>
                <div class="mb-3">
                  <label for="venue" class="form-label">
                    Add New Venue
                  </label>
                  <input
                    type="text"
                    class="form-control"
                    id="venue"
                    value={venueEditText}
                    onChange={onChangeInputEdit}
                  />
                </div>
              </form>
              <button
                type="button"
                class="btn btn-primary btn-sm"
                onClick={onClickAddVenue}
              >
                Add
              </button>
              <hr></hr>
              <form>
                <div class="mb-3">
                  <label for="venue" class="form-label">
                    Delete Venue
                  </label>
                  <input
                    type="text"
                    class="form-control"
                    id="venue"
                    value={venueDeleteText}
                    onChange={onChangeInputDelete}
                  />
                </div>
              </form>
              <button
                type="button"
                class="btn btn-danger btn-sm"
                onClick={onClickDelVenue}
              >
                Remove
              </button>
            </div>
            <div class="modal-footer">
              <button
                type="button"
                class="btn btn-secondary"
                data-bs-dismiss="modal"
              >
                Close
              </button>
            </div>
          </div>
        </div>
      </div>
    </>
  );
}
export default VenueEdit;
