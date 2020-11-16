<html>
	<head>
		<title>Postman for Atlassian Bamboo Server</title>
		<meta name="author" content="Kravchenko E. Dmytro" />
		<meta name="viewport" content="width=device-width, height=device-height, initial-scale=1.0, minimum-scale=1.0" />
		<meta http-equiv="Content-Security-Policy" content="upgrade-insecure-requests" />
		<link rel="stylesheet" type="text/css" href="{{.Config.Web.RootURL}}/resources/bootstrap/css/bootstrap.min.css" />
		<link rel="stylesheet" type="text/css" href="{{.Config.Web.RootURL}}/resources/css/custom/default.css " />

	</head>

	<body>
		<div class="container-fluid">
		
			<div class="row">
				<div class="col">

					<h2>Bamboo Artifact Postman <span class="size-4 badge badge-danger">{{.Version}}</span> </h2>

					<div class="input-group mb-3">
							<input type="text" class="form-control" placeholder="Bamboo Artifact URL" aria-label="Bamboo Artifact URL" aria-describedby="bamboo_artifact_url" name="bamboo_artifact_url" />
							<div class="input-group-append">
								<button class="btn btn-primary" type="button" id="bamboo-postman-lvl-1" data-toggle="tooltip" data-placement="top" title="Send request to server">Send</button>
							</div>
					</div>


				</div>

			</div>

			<div class="row">
				<div class="col">
					Result will be here
				</div>
			</div>

			<div class="border border-primary" id="return_data">
				<div class="row"></div>
			</div>
	
			<div class="border" id="parser_data">
			</div>

		</div>



<!-- Modal -->
<div class="modal fade" id="exampleModal" tabindex="-1" role="dialog" aria-labelledby="exampleModalLabel" aria-hidden="true">
  <div class="modal-dialog" role="document">
    <div class="modal-content">
      <div class="modal-header">
        <h5 class="modal-title" id="exampleModalLabel">In progress...</h5>
        <button type="button" class="close" data-dismiss="modal" aria-label="Close">
          <span aria-hidden="true">&times;</span>
        </button>
      </div>
      <div class="modal-body">
      After complete we close it.
      </div>
    </div>
  </div>
</div>



		<script src="{{.Config.Web.RootURL}}/resources/js/jquery/3.5.1/jquery.min.js"></script>
		<script src="{{.Config.Web.RootURL}}/resources/bootstrap/js/bootstrap.min.js"></script>
		<script src="{{.Config.Web.RootURL}}/resources/js/custom/default.js?buid={{.Config.Web.Build}}"></script>



	</body>

</html>
