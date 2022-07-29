# Configuration
To configure eol-checker, a config.yaml file can be passed with --config. Configuration read in through a file will overwrite the same configuration specified by a flag. If no config file is passed, and no flags are set, reasonable defaults will be used.


Here is a sample of configuration file

```yaml

default:
  url: https://endoflife.date
  alert:
    month: 12
products:
  - name: "go"
    version: 1.17

  - name: "java"
    version: 11


plugins: 
  slack:
    webhook_url: https://hooks.slack.com/services/####/#####
    attachments: 
      - author: 
          name: End Of Life Notifier
          link: https://cloudposse.com/wp-content/uploads/sites/29/2018/02/small-cute-robot-square.png
          icon: https://cloudposse.com/wp-content/uploads/sites/29/2018/02/small-cute-robot-square.png
        color: good
        fallback: Deployed to Staging env
        footer: helm deployments
        footer_icon: https://cloudposse.com/wp-content/uploads/sites/29/2018/02/kubernetes.png
        title: Environment Updated
        title_link: http://demo1.cloudposse.com
        thumb_url: https://cloudposse.com/wp-content/uploads/sites/29/2018/02/SquareLogo2.png 
    user_name: Robot
    icon_emoji: white_check_mark
```