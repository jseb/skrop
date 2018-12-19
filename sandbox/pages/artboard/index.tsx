import React from "react";
import { Grid, Header, Form } from "semantic-ui-react";
import FilterProps from "./filter-props";
import Crop from "models/crop";

export interface IArtboard {
  selectedImageUrl: string;
}

const Artboard: React.SFC<IArtboard> = ({ selectedImageUrl, filters }) => (
  <div>
    <Grid columns={3}>
      <Grid.Row>
        <Grid.Column width="4">
          <Header as="h3">Filters</Header>
        </Grid.Column>
        <Grid.Column>
          <img src={selectedImageUrl} alt="Artboard Image" />{" "}
        </Grid.Column>
        <Grid.Column width="4">
          <Header as="h3">Filters</Header>
          <Form>
            <Form.Group widths="equal">
              <FilterProps filters={[new Crop(1, 1, true)]} />
            </Form.Group>
          </Form>
        </Grid.Column>
      </Grid.Row>
    </Grid>
  </div>
);
export default Artboard;
