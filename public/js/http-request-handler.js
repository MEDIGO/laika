var httpRequestHandler = {
  /**
   * Handles HTTP requests
   * @param {string} method - HTTP method of the request (GET, POST or PATCH)
   * @param {string} url - URL for the request
   * @param {object} info - information to send in the body of the request
   * @param {function} onComplete - function to call after the request is complete
   */
  sendRequest: function(method, url, info, onComplete) {
    var request = new XMLHttpRequest();
    request.open(method, url, true);

    request.onreadystatechange = function (event) {
      if (request.readyState === 4) {
        if (request.status === 200 || request.status === 201) {
          onComplete()
        } else {
           onComplete('Ops! Something went wrong!')
        }
      }
    };

    request.setRequestHeader('Content-Type', 'application/json;charset=UTF-8');
    request.send(JSON.stringify(info));
  }
};
