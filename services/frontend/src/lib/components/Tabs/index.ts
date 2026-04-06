import { Tabs as TabsPrimitive } from "bits-ui";
import Content from "./component-content.svelte";
import List from "./component-list.svelte";
import Trigger from "./component-trigger.svelte";

const Root = TabsPrimitive.Root;

export {
  Root,
  Content,
  List,
  Trigger,
  //
  Root as Tabs,
  Content as TabsContent,
  List as TabsList,
  Trigger as TabsTrigger,
};