var requestOnline = function () {
    $.ajax({
        type: "POST",
        url: "/ajax/online",
        data: JSON.stringify({ online: true }),
    });
};

var interval = 1000 * 5; // request once per 5 seconds

setInterval(requestOnline, interval);