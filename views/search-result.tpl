<!DOCTYPE html>

<body id="body">
	<div class="table-wrapper">
	<table class="t_displaymitra" align="center" width="1000">
	<div class="container-fluid jumbotron" id="jumbotron">
		<center><h1>{{.title}}<h1></center>
	</div>

		<thead align = "center">
			
			<tr>
				<th><center>Nama Film</center></th>
				<th><center>Durasi</center></th>
				<th><center>Tanggal Rilis</center></th>
				<th><center>Rating</center></th>
				<th><center>Sinopsis</center></th>
				
				
			</tr>
		</thead>

			<tr>
				<td align="center">{{.isi1}}</td>
				<td align="center">{{.isi2}}</th>
				<td align="center">{{.isi3}}</th>
				<td align="center">{{.isi4}}</th>
				<td align="center">{{.isi5}}</th>
				
				
				
			</tr>
	</table></div>

	<br><br>
	<div class="container-fluid" id="row">
		<center>
			
			
				<table>
					<tr>
						<td colspan="2" style="padding-top:20px">
							<a href="/"><button type="button" class="btn btn-primary btn-block" style="width:100%">Home</button></a><br>
						</td>
						<td colspan="2" style="padding-top:20px">
							<a href="/searcha"><button type="button" class="btn btn-primary btn-block" style="width:100%">Search Again</button></a><br>
						</td>
					</tr>
				</table>
			
		</center>
	</div>
</body>

