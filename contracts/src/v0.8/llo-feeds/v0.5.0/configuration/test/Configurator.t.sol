import {IConfigurator} from "../interfaces/IConfigurator.sol";
import {Test} from "forge-std/Test.sol";
import {Configurator} from "../Configurator.sol";
import {ExposedConfigurator} from "./mocks/ExposedConfigurator.sol";

/**
 * @title ConfiguratorTest
 * @author samsondav
 * @notice Base class for Configurator tests
 */
contract ConfiguratorTest is Test {
    ExposedConfigurator public configurator;
    // TODO: Write some tests

    function setUp() public virtual {
        channelConfigStore = new ExposedChannelConfigStore();
    }

    function testTypeAndVersion() public view {
        assertEq(channelConfigStore.typeAndVersion(), "ChannelConfigStore 0.0.1");
    }
}