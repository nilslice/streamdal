import { ErrorType } from "../form/validate.ts";
import { AbortCondition } from "streamdal-protos/protos/sp_pipeline.ts";
import { FormBoolean } from "../form/formBoolean.tsx";
import { RadioGroup } from "../form/radioGroup.tsx";
import { MetaData } from "./metadata.tsx";
import IconInfoCircle from "tabler-icons/tsx/info-circle.tsx";
import { Tooltip } from "../tooltip/tooltip.tsx";
import { useEffect, useState } from "preact/hooks";
import { initFlowbite } from "flowbite";
import IconChevronUp from "tabler-icons/tsx/chevron-up.tsx";
import IconChevronDown from "tabler-icons/tsx/chevron-down.tsx";

export type StepConditionType = {
  name: string;
  label: string;
  toolTip: string;
  data: any;
  errors: ErrorType;
};

export const StepCondition = (
  { name, label, toolTip, data, errors }: StepConditionType,
) => {
  const [expanded, setExpanded] = useState(false);

  useEffect(() => {
    initFlowbite();
  }, []);

  return (
    <div
      class={`my-2 border rounded-sm ${
        errors && Object.keys(errors).some((k) => k.includes(name)) &&
        "border-streamdalRed"
      }`}
    >
      <div
        class="flex flex-row justify-between items-center p-2 cursor-pointer"
        onClick={() => setExpanded(!expanded)}
      >
        <div class="flex flex-row justify-start items-center">
          <label
            className={`text-xs ${errors[name] && "text-streamdalRed"}`}
          >
            {label}
          </label>
          <IconInfoCircle
            class="w-4 h-4 ml-1"
            data-tooltip-target={`${name}-tooltip`}
          />
          <Tooltip
            targetId={`${name}-tooltip`}
            message={toolTip}
          />
        </div>
        {expanded
          ? <IconChevronUp class="w-6 h-6 text-twilight" />
          : <IconChevronDown class="w-6 h-6 text-twilight" />}
      </div>
      <div
        className={`flex flex-col p-2 w-full border-t md:${
          expanded ? "visible" : "hidden"
        }`}
      >
        <FormBoolean
          name={`${name}.notify`}
          data={data}
          display="Notify"
        />
        <RadioGroup
          name={`${name}.abort`}
          data={data}
          errors={errors}
          options={{
            [AbortCondition.UNSET]: "Don't Abort",
            [AbortCondition.ABORT_CURRENT]: "Abort Current Pipeline",
            [AbortCondition.ABORT_ALL]: "Abort All Pipelines",
          }}
        />
        <MetaData
          name={`${name}.metadata`}
          data={data}
          errors={errors}
        />
      </div>
    </div>
  );
};
