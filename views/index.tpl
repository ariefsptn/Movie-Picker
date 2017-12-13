<!DOCTYPE html>

<html>
<head>
	<title>{{.title}}</title>
	<meta charset="UTF-8">
	<meta name="viewport" content="width=device-width, initial-scale=1">
	<link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/bootstrap/3.3.7/css/bootstrap.min.css">
	<link rel="stylesheet" type="text/css" href="static/css/customized-style.css">
	<script src="https://maxcdn.bootstrapcdn.com/bootstrap/3.3.7/js/bootstrap.min.css"></script>
</head>

<body id="body">
	
	

	<div class="container-fluid jumbotron" id="jumbotron">
		<center><h1>Movie Picker<h1></center>
	</div>
	
	<div class="container-fluid" id="row">
		<center>
			
			<form class="form-group" action="/search" name="registration-form" method="post">
				<table>
					<tr>
						<td style="padding-right:10px">Minimum Rating</td>
						<td><input type="number" min="1" max="9" name="minimum" class="form-control" required></td>
					</tr>
					<tr>
						<td style="padding-right:10px">Minimum Tahun</td>
						<td><input type="number" min="1940" max="2020" name="mintahun" class="form-control"></td>
					</tr>
					<tr >
						<td style="padding-right:10px">Maximum Tahun</td>
						<td><input type="number" min="1940" max="2020" name="maxtahun" class="form-control"></font></td>
					</tr>
					<tr>
						<td style="padding-right:10px">Genre 1</td>
				
						<td><select name="genre1" class="form-control" required>
					    <option value=""></option>
					    <option value="Action">Action</option>
					    <option value="Comedy">Comedy</option>
					    <option value="Drama">Drama</option>
					    <option value="History">History</option>
					    <option value="Romance">Romance</option>
					    <option value="Horror">Horror</option>
					    <option value="Thriller">Thriller</option>
					    <option value="Adventure">Adventure</option>
					    <option value="Science Fiction">Sci-Fi</option>
					    <option value="Fantasy">Fantasy</option>
					    <option value="Animation">Animation</option>
					    <option value="Crime">Crime</option>
					    <option value="Documentary">Documentary</option>
					    <option value="Foreign">Foreign</option>
					    <option value="Music">Music</option>
					    <option value="Mystery">Mystery</option>
					    <option value="War">War</option>
					    <option value="Western">Western</option>
					</tr>
					<tr>
						<td style="padding-right:10px">Genre 2</td>
							<td><select name="genre2" class="form-control">
					    <option value=""></option>
					    <option value="Action">Action</option>
					    <option value="Comedy">Comedy</option>
					    <option value="Drama">Drama</option>
					    <option value="History">History</option>
					    <option value="Romance">Romance</option>
					    <option value="Horror">Horror</option>
					    <option value="Thriller">Thriller</option>
					    <option value="Adventure">Adventure</option>
					    <option value="Science Fiction">Sci-Fi</option>
					    <option value="Fantasy">Fantasy</option>
					    <option value="Animation">Animation</option>
					    <option value="Crime">Crime</option>
					    <option value="Documentary">Documentary</option>
					    <option value="Foreign">Foreign</option>
					    <option value="Music">Music</option>
					    <option value="Mystery">Mystery</option>
					    <option value="War">War</option>
					    <option value="Western">Western</option>
					</tr>
					<tr>
						<td style="padding-right:10px">Genre 3</td>
							<td><select name="genre3" class="form-control">
					    <option value=""></option>
					    <option value="Action">Action</option>
					    <option value="Comedy">Comedy</option>
					    <option value="Drama">Drama</option>
					    <option value="History">History</option>
					    <option value="Romance">Romance</option>
					    <option value="Horror">Horror</option>
					    <option value="Thriller">Thriller</option>
					    <option value="Adventure">Adventure</option>
					    <option value="Science Fiction">Sci-Fi</option>
					    <option value="Fantasy">Fantasy</option>
					    <option value="Animation">Animation</option>
					    <option value="Crime">Crime</option>
					    <option value="Documentary">Documentary</option>
					    <option value="Foreign">Foreign</option>
					    <option value="Music">Music</option>
					    <option value="Mystery">Mystery</option>
					    <option value="War">War</option>
					    <option value="Western">Western</option>
					</tr>
					<tr>
						<td colspan="2" style="padding-top:20px">
							<input type="submit" value="Pick" class="btn btn-primary btn-block">
						</td>
					</tr>
				</table>
			</form>
		</center>
	</div>

</body>
</html>
