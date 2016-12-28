$(function() {
    if (!window["WebSocket"]) {
        return;
    }

    var content = $("#content");
    var conn = new WebSocket('ws://' + window.location.host + '/ws/itishot');
    var sessionId = null;

    // Textarea is editable only when socket is opened.
    conn.onopen = function(e) {
        content.attr("disabled", false);
    };

    conn.onclose = function(e) {
        content.attr("disabled", true);
    };

    // Whenever we receive a message, update textarea
    conn.onmessage = function(e) {
        if (e.data != content.val()) {
            content.val(e.data);
        }
    };

    var timeoutId = null;
    var typingTimeoutId = null;
    var isTyping = false;

    content.on("keydown", function() {
        isTyping = true;
        window.clearTimeout(typingTimeoutId);
    });

    content.on("keyup", function() {
        typingTimeoutId = window.setTimeout(function() {
            isTyping = false;
        }, 1000);

        window.clearTimeout(timeoutId);
        timeoutId = window.setTimeout(function() {
            if (isTyping) return;
            conn.send(content.val());
        }, 1100);
    });

    $('#submit').on('click', function(){
      if (sessionId === null){
        $.post('/db', JSON.stringify({Content: content.val()}), function(data, err){;
          sessionId = data;
          console.log("Saved document, created id:", sessionId);
        });
      } else {
        $.post('/db', JSON.stringify({Id: sessionId, Content: content.val()}), function(data, err){
          console.log('saved document with session id:', data, err);
        });
      }
    });

    $('#dataz').on('click', function(){
      console.log('dataz clicked');
      $.get('/db', function(data, err){
        console.log('dataz!', data, err);
      });
    });
});
