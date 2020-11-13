<html>
	<head>
		<title>Postman for Atlassian Bamboo Server</title>
		<meta name="author" content="Kravchenko E. Dmytro" />
		<meta name="viewport" content="width=device-width, height=device-height, initial-scale=1.0, minimum-scale=1.0">
		<link rel="stylesheet" type="text/css" href="/resources/bootstrap/css/bootstrap.min.css" />
		<link rel="stylesheet" type="text/css" href="/resources/css/custom/default.css " />
	</head>

	<body>
		<div class="container-fluid">
		
			<div class="row">
				<div class="col">

					<h2>Bamboo Artifact Postman <span class="size-4 badge badge-danger">0.1</span> </h2>

					<div class="input-group mb-3">
							<input type="text" class="form-control" placeholder="Bamboo Artifact URL" aria-label="Bamboo Artifact URL" aria-describedby="bamboo-postman-lvl-1">
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

		</div>

		<script src="/resources/js/jquery/3.5.1/jquery.min.js"></script>
		<script src="/resources/bootstrap/js/bootstrap.min.js"></script>
		<script src="/resources/js/custom/default.js"></script>
	</body>

</html>
