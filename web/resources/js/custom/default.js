$(document).ready(function(){

$('input[name="bamboo_artifact_url"]').val( "" );
var idx = 0
var hostname = "bamboo.macpaw.space/postman"
//var hostname = "localhost"

var clientCard = function ( client ) { 
	tmp = '<address><strong>MacOS Version ' + client.Version+ '</strong><br />';
	tmp += '<footer class="blockquote-footer">Over ' +client.Username+ ' on <cite title="Source Title">'+client.IP+' ['+client.Finger+']</cite></footer>';
	return tmp
}

setID = function( id, content, client ) {

	if (content == null ) {
		return ''
	}
	
	var messages = []
	for ( let i = 0; i < content.length; i++ ) {
		messages[ i ] = content[ i ].Description 
	}

	tpl = '<p><a class="btn btn-primary" data-toggle="collapse" href="#error-'+id+'" role="button" aria-expanded="false" aria-controls="error-'+id+'">Errors</a></p>'
	tpl += '<div class="collapse" id="error-'+id+'"><div class="card card-body">' + client.IP + ':<br />'+messages.join('<br />') + '</div></div>'
	return tpl
}

	$("#bamboo-postman-lvl-1").on("click", function(){
	
		$("#return_data").html('');

		if ( $('input[name="bamboo_artifact_url"]').val() == "" ) {
			alert ( "Please type URL" );
			return
		}

		var socket = new WebSocket("ws://" + hostname + "/gathering", "Upgrade");
		console.log ( "Defined" +  $('input[name="bamboo_artifact_url"]').val() )
		socket.onopen = function(event) {
			 socket.send($('input[name="bamboo_artifact_url"]').val() )
				$('#exampleModal').modal('show');
		};

		socket.onmessage = function(event) {
    	console.log("сообщение от сервера получено", event.data);
			var respond = JSON.parse( event.data );

				if (respond.Status == true) {
					Status = "Finished"
					bg = ""
				} else {
					Status = "Terminated"
					bg = "bg-warning"
				}	
		
				$("#return_data").append ('<div class="row"><div class="col-sm '+bg+'">' + clientCard( respond.Client )+ '</div><div class="col-sm '+bg+'">' + respond.RemoteFile+ '</div><div class="col-sm '+bg+'">' + Status+ '</div><div class="col-sm '+bg+'">'+ setID( idx,  respond.Errors, respond.Client ) +'</div></div>')
		idx++
		};

		socket.onclose = function ( event ) {

			$("#exampleModal").modal( 'toggle' )
		}

	});



})
