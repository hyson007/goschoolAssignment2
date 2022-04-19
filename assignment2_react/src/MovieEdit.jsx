import React, { useState } from "react";
import { toast } from "react-toastify";

function MovieEdit() {
  // const prefix = "http://localhost:8080";
  const prefix = "https://goschoolassignment2.onrender.com";
  const [MovieEditText, setMovieEditText] = useState("");
  const [movieDeleteText, setmovieDeleteText] = useState("");
  const onChangeInputEdit = (e) => {
    setMovieEditText(e.target.value);
  };

  const onChangeInputDelete = (e) => {
    setmovieDeleteText(e.target.value);
  };

  const onClickAddMovie = async (e) => {
    const url = `${prefix}/movies`;
    const result = await fetch(url, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
        Origin: "http://localhost:3000",
      },
      body: JSON.stringify({
        movie: MovieEditText,
      }),
    });
    const tmp = await result.json();
    if (!result.ok) {
      toast.error(tmp.error);
      setMovieEditText("");
      return;
    }
    if (tmp.status === "ok") {
      setMovieEditText("");
      window.location.reload();
    }
    e.preventDefault();
  };

  const onClickDelMovie = async (e) => {
    const url = `${prefix}/movies/${movieDeleteText}`;
    const result = await fetch(url, {
      method: "DELETE",
      headers: {
        Origin: "http://localhost:3000",
      },
    });
    const tmp = await result.json();
    if (!result.ok) {
      toast.error(tmp.error);
      setmovieDeleteText("");
      return;
    }
    if (tmp.status === "ok") {
      setmovieDeleteText("");
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
        data-bs-target="#editMovies"
      >
        Edit Movie
      </button>
      <div class="modal fade" id="editMovies" tabIndex="-1">
        <div class="modal-dialog">
          <div class="modal-content">
            <div class="modal-header">
              <h5 class="modal-title" id="editMoviesLabel">
                Edit Movie
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
                  <label for="movie" class="form-label">
                    Add New Movie
                  </label>
                  <input
                    type="text"
                    class="form-control"
                    id="movie"
                    value={MovieEditText}
                    onChange={onChangeInputEdit}
                  />
                </div>
              </form>
              <button
                type="button"
                class="btn btn-primary btn-sm"
                onClick={onClickAddMovie}
              >
                Add
              </button>
              <hr></hr>
              <form>
                <div class="mb-3">
                  <label for="movie" class="form-label">
                    Delete Movie
                  </label>
                  <input
                    type="text"
                    class="form-control"
                    id="movie"
                    value={movieDeleteText}
                    onChange={onChangeInputDelete}
                  />
                </div>
              </form>
              <button
                type="button"
                class="btn btn-danger btn-sm"
                onClick={onClickDelMovie}
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
export default MovieEdit;
