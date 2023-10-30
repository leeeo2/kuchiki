<template>
     <div class="layout-demo">
          <a-layout style="height: 100%;">
               <a-layout-header>Header</a-layout-header>
               <a-layout>
                    <a-layout-sider theme="dark">Sider</a-layout-sider>
                    <a-layout-content style="height:100px">Content</a-layout-content>
               </a-layout>
               <a-layout-footer>Footer</a-layout-footer>
          </a-layout>
     </div>
</template>

<script>
export default {
     name: "HomePage",
     props: {
          msg: String
     }
}
</script>
