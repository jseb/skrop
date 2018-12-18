import { Form } from "semantic-ui-react";

/**
 *
 * @param filter Instance
 */
export const renderProps = filter =>
  filter.getProps().map(({ prop, type, value }) => {
    switch (type) {
      case "string":
      case "number":
        return (
          <Form.Input
            fluid
            label={prop}
            placeholder={prop}
            key={prop}
            value={value}
          />
        );
        break;
      case "boolean":
        return (
          <Form.Checkbox fluid label={prop} key={prop} defaultChecked={value} />
        );
        break;
      default:
        return null;
    }
  });
