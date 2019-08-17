<template>
  <div>
    <textarea id="mde"></textarea>
  </div>
</template>


<script>
import SimpleMDE from 'simplemde';
import marked from 'marked';

export default {
  name: 'Editor',
  props: {
    value: String,
    name: String,
    previewClass: String,
    autoinit: {
      type: Boolean,
      default() {
        return true;
      },
    },
    highlight: {
      type: Boolean,
      default() {
        return false;
      },
    },
    sanitize: {
      type: Boolean,
      default() {
        return false;
      },
    },
    configs: {
      type: Object,
      default() {
        return {};
      },
    },
  },
  mounted() {
    if (this.autoinit) this.initialize();
  },
  activated() {
    const editor = this.simplemde;
    if (!editor) return;
    const isActive = editor.isSideBySideActive() || editor.isPreviewActive();
    if (isActive) editor.toggleFullScreen();
  },
  methods: {
    initialize() {
      const configs = Object.assign({
        element: document.getElementById('mde'),
        initialValue: this.value,
        renderingConfig: {},
      }, this.configs);
      if (configs.initialValue) {
        this.$emit('input', configs.initialValue);
      }
      if (this.highlight) {
        configs.renderingConfig.codeSyntaxHighlighting = true;
      }
      marked.setOptions({ sanitize: this.sanitize });
      this.simplemde = new SimpleMDE(configs);
      const className = this.previewClass || '';
      this.addPreviewClass(className);
      this.bindingEvents();
    },
    bindingEvents() {
      this.simplemde.codemirror.on('change', () => {
        this.$emit('input', this.simplemde.value());
      });
    },
    addPreviewClass(className) {
      const wrapper = this.simplemde.codemirror.getWrapperElement();
      const preview = document.createElement('div');
      wrapper.nextSibling.className += ` ${className}`;
      preview.className = `editor-preview ${className}`;
      wrapper.appendChild(preview);
    },
  },
  destroyed() {
    this.simplemde = null;
  },
  watch: {
    value(val) {
      if (val === this.simplemde.value()) return;
      this.simplemde.value(val);
    },
  },
};
</script>

<style>
.vue-simplemde .markdown-body {
  padding: 0.5em
}
.vue-simplemde .editor-preview-active, .vue-simplemde .editor-preview-active-side {
  display: block;
}
</style>
