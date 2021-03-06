var options = {
    weekday: "long", year: "numeric", month: "short",
    day: "numeric", hour: "2-digit", minute: "2-digit",
    second: "2-digit"
};

class Clock extends React.Component {
    constructor(props) {
        super(props);
        this.startClock();
        this.state = {
            currentTime: new Date().toLocaleDateString("en-us", options)
        };
    }

    startClock() {
        setInterval(() => {
            this.setState({
                currentTime: new Date().toLocaleDateString("en-us", options)
            });
        }, 1000);
    }

    render() {
        return React.createElement(
            "div",
            null,
            React.createElement(AnalogDisplay, { time: this.state.currentTime }),
            React.createElement(DigitalDisplay, { time: this.state.currentTime })
        );
    }
}