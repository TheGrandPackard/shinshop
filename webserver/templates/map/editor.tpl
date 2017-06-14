{{template "header" .}}
<div class="container">
	<div class="row" style="margin-top:60px;">
		<div class="row">
			<div class="col-md-12">
				<ul class="breadcrumb">
					<li><a href="/map">Map</a></li>
					<li class="active">Editor</li>
				</ul>
			</div>
		</div>
		<div class="col-md-8 col-md-offset-2">
		<div class="row">
				<form class="well bs-component" action="/map/editor/" method="GET">
					<div class="form-group">
						<label class="control-label">MapName</label>
						<div class="input-group">
							<input type="text" class="form-control" name="name" placeholder="{{.LastSearchQuery}}">
							<span class="input-group-btn">
								<button class="btn btn-default" type="button">Find</button>
							</span>
						</div>
					</div>
				</form>
			</div>
			<div class="row">
				<div class="col-md-12">
					<div id="map_container" style="width: 100%; height: 800px; border:1px solid black; ">
						<svg id="map_svg" xmlns="http://www.w3.org/2000/svg" style="display: inline; width: inherit; min-width: inherit; max-width: inherit; height: inherit; min-height: inherit; max-height: inherit;"></svg>
					</div>
				</div>
			</div>
		</div>
	</div>
</div>

<script>var name = {{.LastSearchQuery}}</script>
<script type="text/javascript" src="/js/map/snap.svg.js"></script>
<script type="text/javascript" src="/js/map/svg-pan-zoom.js"></script>
<script type="text/javascript" src="/js/map/map.js"></script>
