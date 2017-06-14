Gb GAE
======

gb gae is a plugin for gb that adds basic project awareness for appengine applications.

It adds the following commands to gb:

	serve		starts a local development App Engine server
	deploy		deploys your application to App Engine
	build		compile packages and dependencies
	test		test packages
	raw		Directly call the dev_appserver.py
	appcfg		Directly call the appcfg.py
	gcloud		Directly call the gcloud command.

goapp, dev_appserver.py etc are automatically wrapped and the correct environment variables are passed to them to make them aware of the project

Contribute
==========

Contributions and suggestions for improvements are very welcome.
I currently don't use managed VMs at all, so any ideas or input on how to make this tool work for them as well are certainly very welcome.
PRs and issues are also welcome and will be checked asap

Legal
=======

Copyright 2015 Palm Stone Games, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
