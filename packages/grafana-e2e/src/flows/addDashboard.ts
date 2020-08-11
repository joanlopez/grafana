import { DashboardTimeRangeConfig, setDashboardTimeRange } from './setDashboardTimeRange';
import { DeleteDashboardConfig } from './deleteDashboard';
import { e2e } from '../index';
import { getDashboardUid } from '../support/url';
import { selectOption } from './selectOption';

export interface AddDashboardConfig {
  timeRange: DashboardTimeRangeConfig;
  timezone: string;
  title: string;
  variables: AddVariableConfig[];
}

export interface AddVariableConfig {
  dataSource?: string;
  hide?: string;
  label: string;
  name: string;
  query?: string;
  regex?: string;
  type: string;
}

// @todo this actually returns type `Cypress.Chainable`
export const addDashboard = (config?: Partial<AddDashboardConfig>): any => {
  const fullConfig = {
    timeRange: {
      from: '2020-01-01 00:00:00',
      to: '2020-01-01 06:00:00',
    },
    timezone: 'Coordinated Universal Time',
    title: `e2e-${Date.now()}`,
    variables: [],
    ...config,
  } as AddDashboardConfig;

  const { timeRange, timezone, title, variables } = fullConfig;

  e2e().logToConsole('Adding dashboard with title:', title);

  e2e.pages.AddDashboard.visit();

  e2e.pages.Dashboard.Toolbar.toolbarItems('Dashboard settings').click();

  // @todo use the time range picker's time zone control
  selectOption(e2e.pages.Dashboard.Settings.General.timezone(), timezone);

  addVariables(variables);

  e2e.components.BackButton.backArrow().click();

  setDashboardTimeRange(timeRange);

  e2e.pages.Dashboard.Toolbar.toolbarItems('Save dashboard').click();

  e2e.pages.SaveDashboardAsModal.newName()
    .clear()
    .type(title);
  e2e.pages.SaveDashboardAsModal.save().click();

  e2e.flows.assertSuccessNotification();

  e2e().logToConsole('Added dashboard with title:', title);

  return e2e()
    .url()
    .then((url: string) => {
      const uid = getDashboardUid(url);

      e2e.getScenarioContext().then(({ addedDashboards }: any) => {
        e2e.setScenarioContext({
          addedDashboards: [...addedDashboards, { title, uid } as DeleteDashboardConfig],
        });
      });

      // @todo remove `wrap` when possible
      return e2e().wrap({
        config: fullConfig,
        uid,
      });
    });
};

const addVariable = (config: Partial<AddVariableConfig>, first: boolean): any => {
  const fullConfig = {
    type: 'Query',
    ...config,
  } as AddVariableConfig;

  if (first) {
    e2e.pages.Dashboard.Settings.Variables.List.addVariableCTA().click();
  } else {
    e2e.pages.Dashboard.Settings.Variables.List.newButton().click();
  }

  const { dataSource, hide, label, name, query, regex, type } = fullConfig;

  if (hide) {
    e2e.pages.Dashboard.Settings.Variables.Edit.General.generalHideSelect().select(hide);
  }

  e2e.pages.Dashboard.Settings.Variables.Edit.General.generalLabelInput().type(label);
  e2e.pages.Dashboard.Settings.Variables.Edit.General.generalNameInput().type(name);

  if (type !== 'Query') {
    e2e.pages.Dashboard.Settings.Variables.Edit.General.generalTypeSelect().select(type);
  }

  if (dataSource && (type === 'Ad hoc filters' || type === 'Datasource' || type === 'Query')) {
    e2e.pages.Dashboard.Settings.Variables.Edit.QueryVariable.queryOptionsDataSourceSelect().select(dataSource);
  }

  if (query && type === 'Query') {
    e2e.pages.Dashboard.Settings.Variables.Edit.QueryVariable.queryOptionsQueryInput().type(query);
  }

  if (regex && type === 'Query') {
    e2e.pages.Dashboard.Settings.Variables.Edit.QueryVariable.queryOptionsRegExInput().type(regex);
  }

  e2e.pages.Dashboard.Settings.Variables.Edit.General.addButton().click();

  return fullConfig;
};

const addVariables = (configs: AddVariableConfig[]): any => {
  if (configs.length > 0) {
    e2e.pages.Dashboard.Settings.General.sectionItems('Variables').click();
  }

  return configs.map((config, i) => addVariable(config, i === 0));
};
