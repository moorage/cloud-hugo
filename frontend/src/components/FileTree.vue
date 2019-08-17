<template>
<ul id="file-tree">
  <tree-item
    class="item"
    :item="treeData"
    @make-folder="makeFolder"
    @add-item="addItem"
  ></tree-item>
</ul>
</template>
<script>
export default {
  el: '#file-tree',
  data: {
    treeData: treeData
  },
  methods: {
  	makeFolder: function (item) {
    	Vue.set(item, 'children', [])
      this.addItem(item)
    },
    addItem: function (item) {
    	item.children.push({
        name: 'new stuff'
      })
    }
  }
}
</script>

<style scoped>
body {
  font-family: Menlo, Consolas, monospace;
  color: #444;
}
.item {
  cursor: pointer;
}
.bold {
  font-weight: bold;
}
ul {
  padding-left: 1em;
  line-height: 1.5em;
  list-style-type: dot;
}
</style>