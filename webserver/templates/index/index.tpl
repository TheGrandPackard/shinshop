{{template "header" .}}
  <div class="container">

	<div class="row" style="margin-top:60px;">
	  	<div class="col-md-8 col-md-offset-2">
	  		<div class="row">
				<form class="well bs-component" action="/character/search/" method="GET">
					<div class="form-group">
			          <label class="control-label">Account Manager</label>
			          <div class="input-group">
			            <input type="text" class="form-control" name="name" placeholder="{{.LastSearchQuery}}">
			            <span class="input-group-btn">
			              <button class="btn btn-default" type="button">Search</button>
			            </span>
			          </div>
			        </div>
			     </form>
			 </div>
			<div class="row">
				
			</div>
		</div>
	</div>
  </div>