function searchStatisticInBar(e) {
  var search = document.getElementById("search").value;
  if (search == "") {
    alert("Please enter the keyword");
    e.preventDefault();
    return;
  }

  alert("Wait a minute...");
  e.preventDefault();

  fetch(
    "/covid/search/bar?" +
      new URLSearchParams({
        country: search,
      }),
    {
      method: "GET",
      headers: {
        "Content-Type": "application/json",
      },
    }
  )
    .then(function (response) {
      // The API call was successful!
      return response.text();
    })
    .then(function (html) {
      document.getElementById("chart").srcdoc = html;
    })
    .catch(function (err) {
      // There was an error
      console.warn("Something went wrong", err);
    });
}

function searchStatisticInBarDefault() {
  fetch(
    "/covid/search/bar?" +
      new URLSearchParams({
        country: "Japan",
      }),
    {
      method: "GET",
      headers: {
        "Content-Type": "application/json",
      },
    }
  )
    .then(function (response) {
      // The API call was successful!
      return response.text();
    })
    .then(function (html) {
      document.getElementById("chart").srcdoc = html;
    })
    .catch(function (err) {
      // There was an error
      console.warn("Something went wrong", err);
    });
}
