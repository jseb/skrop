import React from "react";
import { Tab } from "semantic-ui-react";
import { renderProps } from "ui-components/components-factory/components-factory";

interface IFilterProps {
  filters: object[];
}

const FilterProps: React.SFC<IFilterProps> = ({ filters }) => {
  const filtersProps =
    filters && filters.length ? filters.map(renderProps) : [];
  const panes = [
    {
      menuItem: "Properties",
      render: () => (
        <Tab.Pane>
          <>{filtersProps}</>
        </Tab.Pane>
      )
    },
    {
      menuItem: "Export",
      render: () => <Tab.Pane>{JSON.stringify(filters[0].toJson())}</Tab.Pane>
    }
  ];

  return <Tab panes={panes} />;
};
export default FilterProps;
