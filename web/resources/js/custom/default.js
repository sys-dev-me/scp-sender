$(document).ready(function(){


var idx = 0
var hostname = "116.203.182.159"

setID = function( id, content, client ) {

	if (content == null ) {
		return ''
	}
	
	var messages = []
	for ( let i = 0; i < content.length; i++ ) {
		messages[ i ] = content[ i ].Description 
	}

	tpl = '<p><a class="btn btn-primary" data-toggle="collapse" href="#error-'+id+'" role="button" aria-expanded="false" aria-controls="error-'+id+'">Errors</a></p>'
	tpl += '<div class="collapse" id="error-'+id+'"><div class="card card-body">' + client + ':<br />'+messages.join('<br />') + '</div></div>'
	return tpl
}

	$("#bamboo-postman-lvl-1").on("click", function(){
		var socket = new WebSocket("ws://" + hostname + ":3001/gathering", "Upgrade");

		socket.onopen = function(event) {
    	socket.send("соединение с сервером установлено");
			 socket.send("Lets Gathering")
		};

		socket.onmessage = function(event) {
    	console.log("сообщение от сервера получено", event.data);
			var respond = JSON.parse( event.data );

				if (respond.Status == true) {
					Status = "Finished"
				} else {
					Status = "Terminated"
				}	
		
				$("#return_data").append ('<div class="row"><div class="col-sm">' + respond.Client + '</div><div class="col-sm">' + respond.Filename+ ' -> </div><div class="col-sm">' + respond.RemoteFile+ '</div><div class="col-sm">' + Status+ '</div><div class="col-sm">'+ setID( idx,  respond.Errors, respond.Client ) +'</div></div>')
		idx++
		};

	});
})
