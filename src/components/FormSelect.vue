<template>
  <FormGroup :label="label">
    <select :class="$style.component">
      <option
        v-for="option in options"
        :key="option.value"
        value="val">
        {{ option.label }}
      </option>
    </select>
  </FormGroup>
</template>

<script>
import FormGroup from '@/components/FormGroup';

export default {
  name: 'FormSelect',
  props: {
    label: String,
    options: Array,
  },
  components: {
    FormGroup,
  },
};
</script>

<style lang="scss" module>
.component {
  display: block;
  width: 100%;
  background-color: white;
}
</style>
