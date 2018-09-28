package main

var goMod = `module %v`

var mainGo = `package main

import "fmt"

func main() {
	fmt.Println("Hello World!")
}
`

var iml = `<?xml version="1.0" encoding="UTF-8"?>
<module type="WEB_MODULE" version="4">
  <component name="NewModuleRootManager">
    <content url="file://$MODULE_DIR$" />
    <orderEntry type="inheritedJdk" />
    <orderEntry type="sourceFolder" forTests="false" />
  </component>
</module>
`

var modules = `<?xml version="1.0" encoding="UTF-8"?>
<project version="4">
  <component name="ProjectModuleManager">
    <modules>
      <module fileurl="file://$PROJECT_DIR$/.idea/%v.iml" filepath="$PROJECT_DIR$/.idea/%v.iml" />
    </modules>
  </component>
</project>
`

var watcherTasks = `<?xml version="1.0" encoding="UTF-8"?>
<project version="4">
  <component name="ProjectTasksOptions">
    <TaskOptions isEnabled="true">
      <option name="arguments" value="-w $FilePath$" />
      <option name="checkSyntaxErrors" value="true" />
      <option name="description" />
      <option name="exitCodeBehavior" value="ERROR" />
      <option name="fileExtension" value="go" />
      <option name="immediateSync" value="false" />
      <option name="name" value="goimports" />
      <option name="output" value="$FilePath$" />
      <option name="outputFilters">
        <array />
      </option>
      <option name="outputFromStdout" value="false" />
      <option name="program" value="goimports" />
      <option name="runOnExternalChanges" value="false" />
      <option name="scopeName" value="Project Files" />
      <option name="trackOnlyRoot" value="true" />
      <option name="workingDir" value="" />
      <envs>
        <env name="GOROOT" value="$GOROOT$" />
        <env name="GOPATH" value="$GOPATH$" />
        <env name="PATH" value="$GoBinDirs$" />
      </envs>
    </TaskOptions>
  </component>
</project>
`

var workspace = `<?xml version="1.0" encoding="UTF-8"?>
<project version="4">
  <component name="ChangeListManager">
    <list default="true" id="dda3e0b1-56df-47d3-b003-647e6988ab90" name="Default" comment="" />
    <option name="EXCLUDED_CONVERTED_TO_IGNORED" value="true" />
    <option name="TRACKING_ENABLED" value="true" />
    <option name="SHOW_DIALOG" value="false" />
    <option name="HIGHLIGHT_CONFLICTS" value="true" />
    <option name="HIGHLIGHT_NON_ACTIVE_CHANGELIST" value="false" />
    <option name="LAST_RESOLUTION" value="IGNORE" />
  </component>
  <component name="GOROOT" path="%v" />
  <component name="JsBuildToolGruntFileManager" detection-done="true" sorting="DEFINITION_ORDER" />
  <component name="JsBuildToolPackageJson" detection-done="true" sorting="DEFINITION_ORDER" />
  <component name="JsGulpfileManager">
    <detection-done>true</detection-done>
    <sorting>DEFINITION_ORDER</sorting>
  </component>
  <component name="ProjectFrameBounds" fullScreen="true">
    <option name="width" value="1920" />
    <option name="height" value="1080" />
  </component>
  <component name="ProjectView">
    <navigator currentView="ProjectPane" proportions="" version="1">
      <flattenPackages />
      <showMembers />
      <showModules />
      <showLibraryContents />
      <hideEmptyPackages />
      <abbreviatePackageNames />
      <autoscrollToSource />
      <autoscrollFromSource />
      <sortByType />
      <manualOrder />
      <foldersAlwaysOnTop value="true" />
    </navigator>
    <panes>
      <pane id="ProjectPane">
        <subPane>
          <expand>
            <path>
              <item name="%v" type="b2602c69:ProjectViewProjectNode" />
              <item name="%v" type="462c0819:PsiDirectoryNode" />
            </path>
          </expand>
          <select />
        </subPane>
      </pane>
      <pane id="Scratches" />
      <pane id="Scope" />
    </panes>
  </component>
  <component name="PropertiesComponent">
    <property name="settings.editor.selected.configurable" value="watcher.settings" />
    <property name="last_opened_file_path" value="$PROJECT_DIR$" />
    <property name="go.sdk.automatically.set" value="true" />
  </component>
  <component name="RunDashboard">
    <option name="ruleStates">
      <list>
        <RuleState>
          <option name="name" value="ConfigurationTypeDashboardGroupingRule" />
        </RuleState>
        <RuleState>
          <option name="name" value="StatusDashboardGroupingRule" />
        </RuleState>
      </list>
    </option>
  </component>
  <component name="RunManager">
    <configuration name="%v" type="GoApplicationRunConfiguration" factoryName="Go Application" singleton="true">
      <module name="%v" />
      <working_directory value="$PROJECT_DIR$/" />
      <go_parameters value="-i" />
      <kind value="DIRECTORY" />
      <filePath value="$PROJECT_DIR$/" />
      <directory value="$PROJECT_DIR$/" />
    </configuration>
  </component>
  <component name="ShelveChangesManager" show_recycled="false">
    <option name="remove_strategy" value="false" />
  </component>
  <component name="ToolWindowManager">
    <frame x="0" y="0" width="1920" height="1080" extended-state="0" />
    <layout>
      <window_info id="Project" active="true" anchor="left" auto_hide="false" internal_type="DOCKED" type="DOCKED" visible="true" show_stripe_button="true" weight="0.20926517" sideWeight="0.5" order="0" side_tool="false" content_ui="combo" />
      <window_info id="TODO" active="false" anchor="bottom" auto_hide="false" internal_type="DOCKED" type="DOCKED" visible="false" show_stripe_button="true" weight="0.33" sideWeight="0.5" order="6" side_tool="false" content_ui="tabs" />
      <window_info id="Event Log" active="false" anchor="bottom" auto_hide="false" internal_type="DOCKED" type="DOCKED" visible="false" show_stripe_button="true" weight="0.33" sideWeight="0.5" order="-1" side_tool="true" content_ui="tabs" />
      <window_info id="Database" active="false" anchor="right" auto_hide="false" internal_type="DOCKED" type="DOCKED" visible="false" show_stripe_button="true" weight="0.33" sideWeight="0.5" order="-1" side_tool="false" content_ui="tabs" />
      <window_info id="Find" active="false" anchor="bottom" auto_hide="false" internal_type="DOCKED" type="DOCKED" visible="false" show_stripe_button="true" weight="0.33" sideWeight="0.5" order="1" side_tool="false" content_ui="tabs" />
      <window_info id="Run" active="false" anchor="bottom" auto_hide="false" internal_type="DOCKED" type="DOCKED" visible="false" show_stripe_button="true" weight="0.33" sideWeight="0.5" order="2" side_tool="false" content_ui="tabs" />
      <window_info id="Version Control" active="false" anchor="bottom" auto_hide="false" internal_type="DOCKED" type="DOCKED" visible="false" show_stripe_button="false" weight="0.33" sideWeight="0.5" order="-1" side_tool="false" content_ui="tabs" />
      <window_info id="Structure" active="false" anchor="left" auto_hide="false" internal_type="DOCKED" type="DOCKED" visible="false" show_stripe_button="true" weight="0.25" sideWeight="0.5" order="1" side_tool="false" content_ui="tabs" />
      <window_info id="Terminal" active="false" anchor="bottom" auto_hide="false" internal_type="DOCKED" type="DOCKED" visible="false" show_stripe_button="true" weight="0.33" sideWeight="0.5" order="-1" side_tool="false" content_ui="tabs" />
      <window_info id="Debug" active="false" anchor="bottom" auto_hide="false" internal_type="DOCKED" type="DOCKED" visible="false" show_stripe_button="true" weight="0.4" sideWeight="0.5" order="3" side_tool="false" content_ui="tabs" />
      <window_info id="Favorites" active="false" anchor="left" auto_hide="false" internal_type="DOCKED" type="DOCKED" visible="false" show_stripe_button="true" weight="0.33" sideWeight="0.5" order="-1" side_tool="true" content_ui="tabs" />
      <window_info id="Cvs" active="false" anchor="bottom" auto_hide="false" internal_type="DOCKED" type="DOCKED" visible="false" show_stripe_button="true" weight="0.25" sideWeight="0.5" order="4" side_tool="false" content_ui="tabs" />
      <window_info id="Hierarchy" active="false" anchor="right" auto_hide="false" internal_type="DOCKED" type="DOCKED" visible="false" show_stripe_button="true" weight="0.25" sideWeight="0.5" order="2" side_tool="false" content_ui="combo" />
      <window_info id="Message" active="false" anchor="bottom" auto_hide="false" internal_type="DOCKED" type="DOCKED" visible="false" show_stripe_button="true" weight="0.33" sideWeight="0.5" order="0" side_tool="false" content_ui="tabs" />
      <window_info id="Commander" active="false" anchor="right" auto_hide="false" internal_type="DOCKED" type="DOCKED" visible="false" show_stripe_button="true" weight="0.4" sideWeight="0.5" order="0" side_tool="false" content_ui="tabs" />
      <window_info id="Inspection" active="false" anchor="bottom" auto_hide="false" internal_type="DOCKED" type="DOCKED" visible="false" show_stripe_button="true" weight="0.4" sideWeight="0.5" order="5" side_tool="false" content_ui="tabs" />
      <window_info id="Ant Build" active="false" anchor="right" auto_hide="false" internal_type="DOCKED" type="DOCKED" visible="false" show_stripe_button="true" weight="0.25" sideWeight="0.5" order="1" side_tool="false" content_ui="tabs" />
    </layout>
  </component>
  <component name="TypeScriptGeneratedFilesManager">
    <option name="version" value="1" />
  </component>
  <component name="VcsContentAnnotationSettings">
    <option name="myLimit" value="2678400000" />
  </component>
  <component name="XDebuggerManager">
    <breakpoint-manager />
    <watches-manager />
  </component>
</project>
`
