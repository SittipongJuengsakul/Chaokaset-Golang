package app

import "github.com/revel/revel"

func init() {
	// Filters is the default set of global filters.
	revel.Filters = []revel.Filter{
		revel.PanicFilter,             // Recover from panics and display an error page instead.
		revel.RouterFilter,            // Use the routing table to select the right Action
		revel.FilterConfiguringFilter, // A hook for adding or removing per-Action filters.
		//CORSFilter,                    // Cross Origin Resource Sharing
		revel.ParamsFilter,            // Parse parameters into Controller.Params.
		revel.SessionFilter,           // Restore and write the session cookie.
		revel.FlashFilter,             // Restore and write the flash cookie.
		revel.ValidationFilter,        // Restore kept validation errors and save new ones from cookie.
		revel.I18nFilter,              // Resolve the requested language
		CORSFilter,                  // Add some security based headers
		revel.InterceptorFilter,       // Run interceptors around the action.
		revel.CompressFilter,          // Compress the result.
		revel.ActionInvoker,           // Invoke the action.
	}

	// register startup functions with OnAppStart
	// ( order dependent )
	// revel.OnAppStart(InitDB)
	// revel.OnAppStart(FillCache)
}

// TODO turn this into revel.HeaderFilter
// should probably also have a filter for CSRF
// not sure if it can go in the same filter or not
var CORSFilter = func(c *revel.Controller, fc []revel.Filter) {
        c.Response.Out.Header().Set("Access-Control-Allow-Origin", "*")
        c.Response.Out.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
        c.Response.Out.Header().Set("Access-Control-Allow-Headers",
                "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

        // Stop here for a Preflighted OPTIONS request.
        if c.Request.Method == "OPTIONS" {
                return
        }

        fc[0](c, fc[1:]) // Execute the next filter stage.
}
