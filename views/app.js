function getStatisticInBar() {
    fetch('/covid/bar', {
      method: 'GET',
      headers: {
        'Content-Type': 'application/json'
      }
    }).then(function (response) {
      // The API call was successful!
      return response.text();
    }).then(function (html) {
      document.getElementById("chart").srcdoc = html;  
    
    }).catch(function (err) {
      // There was an error
      console.warn('Something went wrong.', err);
    });
}