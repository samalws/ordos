const serverUrl = "192.168.1.14"

function addToChat(speaker,msg) {
  document.getElementById("chatbar").innerHTML += '<p class="chat"><b class="' +
    (speaker == "Me" ? "" : "not") + 'me">'+speaker+'</b>: '+msg+'</p>'
}

function eventRecieved(data,status) {
  console.log(data,status)
  $.get(serverUrl+":5252",eventRecieved) //server waits until there's new info to give then gives it
}

$(document).ready(function() {
  document.getElementById("chatbox").onkeyup = function(key) {
    if (key.key=="Enter") {
      var str = document.getElementById("chatbox").value
      document.getElementById("chatbox").value = ""
      addToChat("Me",str)
      $.post(serverUrl+":6565",str,function(){})
    }
  }

  for (var i=0;i<3;i++) {
    $.get(serverUrl+":5252",eventRecieved) //add 3 event listeners for good measure
  }
})
